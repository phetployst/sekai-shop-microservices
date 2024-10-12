package itemUsecase

import (
	"context"
	"errors"

	"github.com/phetployst/sekai-shop-microservices/modules/item"
	"github.com/phetployst/sekai-shop-microservices/modules/item/itemRepository"
	"github.com/phetployst/sekai-shop-microservices/pkg/utils"
)

type (
	ItemUsecaseService interface {
		CreateItem(pctx context.Context, req *item.CreateItemReq) (string, error)
	}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemRepository}
}

func (u *itemUsecase) CreateItem(pctx context.Context, req *item.CreateItemReq) (string, error) {
	if !u.itemRepository.IsUniqueItem(pctx, req.Title) {
		return "", errors.New("error: this title is already exist")
	}

	itemId, err := u.itemRepository.InsertOneItem(pctx, &item.Item{
		Title:       req.Title,
		Price:       req.Price,
		Damage:      req.Damage,
		UsageStatus: true,
		ImageUrl:    req.ImageUrl,
		CreatedAt:   utils.LocalTime(),
		UpdatedAt:   utils.LocalTime(),
	})
	if err != nil {
		return "", err
	}

	return itemId.String(), err
}
