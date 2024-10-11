package server

import (
	"log"

	"github.com/phetployst/sekai-shop-microservices/modules/player/playerHandler"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerPb"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerRepository"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerUsecase"
	"github.com/phetployst/sekai-shop-microservices/pkg/grpccon"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	player := s.app.Group("/player_v1")

	player.GET("/", s.healthCheckService)
	player.POST("/player/register", httpHandler.CreatePlayer)
	player.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
	player.POST("/player/add-money", httpHandler.AddPlayerMoney)
	player.GET("/player/saving-account/:player_id", httpHandler.GetPlayerSavingAccount)
}
