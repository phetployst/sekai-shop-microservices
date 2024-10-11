package authRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/phetployst/sekai-shop-microservices/config"
	"github.com/phetployst/sekai-shop-microservices/modules/auth"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerPb"
	"github.com/phetployst/sekai-shop-microservices/pkg/grpccon"
	"github.com/phetployst/sekai-shop-microservices/pkg/jwtauth"
	"github.com/phetployst/sekai-shop-microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface {
		CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error)
		InsertOnePlayerCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error)
		FindOnePlayerCredential(pctx context.Context, credentialId string) (*auth.Credential, error)
		FindOnePlayerProfileToRefresh(pctx context.Context, grpcUrl string, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error)
		UpdateOnePlayerCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error
		AccessToken(cfg *config.Config, claims *jwtauth.Claims) string
		RefreshToken(cfg *config.Config, claims *jwtauth.Claims) string
	}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db: db}
}

func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}

func (r *authRepository) CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Player().CredentialSearch(ctx, req)
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %s", err.Error())
		return nil, errors.New("error: email or password is incorrect")
	}

	return result, nil
}

func (r *authRepository) FindOnePlayerProfileToRefresh(pctx context.Context, grpcUrl string, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Player().FindOnePlayerProfileToRefresh(ctx, req)
	if err != nil {
		log.Printf("Error: FindOnePlayerProfileToRefresh failed: %s", err.Error())
		return nil, errors.New("error: player profile not found")
	}

	return result, nil
}

func (r *authRepository) InsertOnePlayerCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	req.CreatedAt = utils.LocalTime()
	req.UpdatedAt = utils.LocalTime()

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayerCredential failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one player credential failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *authRepository) FindOnePlayerCredential(pctx context.Context, credentialId string) (*auth.Credential, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result := new(auth.Credential)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(credentialId)}).Decode(result); err != nil {
		log.Printf("Error: FindOnePlayerCredential failed: %s", err.Error())
		return nil, errors.New("error: find one player credential failed")
	}

	return result, nil
}

func (r *authRepository) UpdateOnePlayerCredential(pctx context.Context, credentialId string, req *auth.UpdateRefreshTokenReq) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	_, err := col.UpdateOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(credentialId)},
		bson.M{
			"$set": bson.M{
				"player_id":     req.PlayerId,
				"access_token":  req.AccessToken,
				"refresh_token": req.RefreshToken,
				"updated_at":    req.UpdatedAt,
			},
		},
	)
	if err != nil {
		log.Printf("Error: UpdateOnePlayerCredential failed: %s", err.Error())
		return errors.New("error: player credential not found")
	}

	return nil
}

func (r *authRepository) AccessToken(cfg *config.Config, claims *jwtauth.Claims) string {
	return jwtauth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtauth.Claims{
		PlayerId: claims.PlayerId,
		RoleCode: int(claims.RoleCode),
	}).SignToken()
}

func (r *authRepository) RefreshToken(cfg *config.Config, claims *jwtauth.Claims) string {
	return jwtauth.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &jwtauth.Claims{
		PlayerId: claims.PlayerId,
		RoleCode: int(claims.RoleCode),
	}).SignToken()
}
