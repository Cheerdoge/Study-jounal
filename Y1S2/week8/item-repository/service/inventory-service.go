package service

import (
	"item-repository/model"
)

type InventoryRepository interface {
	GetItemByName(name string) (model.Item, error)
	GetItemByID(id uint) (model.Item, error)
	GetItems() ([]model.Item, error)
	ReduceItemNum(name string, num int) error
	RepositoryInit() error
	AddProcessedOrder(orderID uint) error
	IsOrderProcessed(orderID uint) (bool, error)
	ProcessOrder(orderID uint, itemname string, nums int) error
}

type ItemService struct {
	Itemrepo InventoryRepository
}

func NewItemService(itemRepo InventoryRepository) *ItemService {
	return &ItemService{
		Itemrepo: itemRepo,
	}
}

func (s *ItemService) GetItemByName(name string) (model.Item, error) {
	return s.Itemrepo.GetItemByName(name)
}

func (s *ItemService) GetItemByID(id uint) (model.Item, error) {
	return s.Itemrepo.GetItemByID(id)
}

func (s *ItemService) GetItems() ([]model.Item, error) {
	return s.Itemrepo.GetItems()
}

func (s *ItemService) ProcessOrder(orderID uint, itemname string, nums int) error {
	return s.Itemrepo.ProcessOrder(orderID, itemname, nums)
}
