package itemUsecase

import "github.com/phetployst/sekai-shop-microservices/modules/item/itemRepository"

type (
	ItemUsecaseService interface{}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemRepository}
}
