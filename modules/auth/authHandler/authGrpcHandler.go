package authHandler

import "github.com/phetployst/sekai-shop-microservices/modules/auth/authUsecase"

type authGrpcHandler struct {
	authUsecase authUsecase.AuthUsecaseService
}

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService {
	return &authGrpcHandler{authUsecase: authUsecase}
}
