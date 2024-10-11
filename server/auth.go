package server

import (
	"log"

	"github.com/phetployst/sekai-shop-microservices/modules/auth/authHandler"
	"github.com/phetployst/sekai-shop-microservices/modules/auth/authPb"
	"github.com/phetployst/sekai-shop-microservices/modules/auth/authRepository"
	"github.com/phetployst/sekai-shop-microservices/modules/auth/authUsecase"
	"github.com/phetployst/sekai-shop-microservices/pkg/grpccon"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	auth.GET("/", s.healthCheckService)
	auth.POST("/auth/login", httpHandler.Login)
}
