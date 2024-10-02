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

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService {
	return &authGrpcHandler{
		authUsecase: authUsecase,
	}
}

func (g *authGrpcHandler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (g *authGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
