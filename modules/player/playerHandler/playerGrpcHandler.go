package playerHandler

import (
	"context"

	"github.com/phetployst/sekai-shop-microservices/modules/player/playerPb"
	"github.com/phetployst/sekai-shop-microservices/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{
		playerUsecase: playerUsecase,
	}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return g.playerUsecase.FindOnePlayerCredential(ctx, req.Password, req.Email)
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}
