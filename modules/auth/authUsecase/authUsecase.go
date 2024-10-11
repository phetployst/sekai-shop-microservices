package authUsecase

import (
	"context"
	"time"

	"github.com/phetployst/sekai-shop-microservices/config"
	"github.com/phetployst/sekai-shop-microservices/modules/auth"
	"github.com/phetployst/sekai-shop-microservices/modules/auth/authRepository"
	"github.com/phetployst/sekai-shop-microservices/modules/player"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerPb"
	"github.com/phetployst/sekai-shop-microservices/pkg/jwtauth"
	"github.com/phetployst/sekai-shop-microservices/pkg/utils"
)

type (
	AuthUsecaseService interface {
		Login(pctx context.Context, cfg *config.Config, req *auth.PlayerLoginReq) (*auth.ProfileIntercepter, error)
	}

	authUsecase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository: authRepository}
}

func (u *authUsecase) Login(pctx context.Context, cfg *config.Config, req *auth.PlayerLoginReq) (*auth.ProfileIntercepter, error) {
	profile, err := u.authRepository.CredentialSearch(pctx, cfg.Grpc.PlayerUrl, &playerPb.CredentialSearchReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	profile.Id = "player:" + profile.Id

	accessToken := u.authRepository.AccessToken(cfg, &jwtauth.Claims{
		PlayerId: profile.Id,
		RoleCode: int(profile.RoleCode),
	})

	refreshToken := u.authRepository.RefreshToken(cfg, &jwtauth.Claims{
		PlayerId: profile.Id,
		RoleCode: int(profile.RoleCode),
	})

	credentialId, err := u.authRepository.InsertOnePlayerCredential(pctx, &auth.Credential{
		PlayerId:     profile.Id,
		RoleCode:     int(profile.RoleCode),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	credential, err := u.authRepository.FindOnePlayerCredential(pctx, credentialId.Hex())
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")

	return &auth.ProfileIntercepter{
		PlayerProfile: &player.PlayerProfile{
			Id:        profile.Id,
			Email:     profile.Email,
			Username:  profile.Username,
			CreatedAt: utils.ConvertStringTimeToTime(profile.CreatedAt).In(loc),
			UpdatedAt: utils.ConvertStringTimeToTime(profile.UpdatedAt).In(loc),
		},
		Credential: &auth.CredentialRes{
			Id:           credential.Id.Hex(),
			PlayerId:     credential.PlayerId,
			RoleCode:     credential.RoleCode,
			AccessToken:  credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreatedAt:    credential.CreatedAt.In(loc),
			UpdatedAt:    credential.UpdatedAt.In(loc),
		},
	}, nil
}
