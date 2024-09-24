package server

import (
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerHandler"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerRepository"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	player := s.app.Group("/player_v1")
	_ = player

}
