package authHandler

import (
	"context"

	"github.com/phetployst/sekai-shop-microservices/modules/auth/authPb"
	"github.com/phetployst/sekai-shop-microservices/modules/auth/authUsecase"
)

type authGrpcHandler struct {
	authPb.UnimplementedAuthGrpcServiceServer
	authUsecase authUsecase.AuthUsecaseService
}

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{
		authUsecase: authUsecase,
	}
}

func (g *authGrpcHandler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return g.authUsecase.AccessTokenSearch(ctx, req.AccessToken)
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return g.authUsecase.RolesCount(ctx)
}
