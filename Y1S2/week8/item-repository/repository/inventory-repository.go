package repository

import (
	"errors"
	"item-repository/model"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	r := &InventoryRepository{db: db}
	err := r.RepositoryInit() // 初始化库存数据
	if err != nil {
		panic(err)
	}
	return r
}

func (r *InventoryRepository) GetItemByName(name string) (model.Item, error) {
	var item model.Item
	if err := r.db.Where("name = ?", name).First(&item).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (r *InventoryRepository) GetItemByID(id uint) (model.Item, error) {
	var item model.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (r *InventoryRepository) GetItems() ([]model.Item, error) {
	var items []model.Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *InventoryRepository) ReduceItemNum(name string, num int) error {
	var item model.Item
	item, err := r.GetItemByName(name)
	if err != nil {
		return err
	}
	item.Num -= num
	if err := r.db.Save(item).Error; err != nil {
		return err
	}
	return nil
}

func (r *InventoryRepository) RepositoryInit() error {
	items := []model.Item{
		{Name: "item1", Num: 0},
		{Name: "item2", Num: 200},
		{Name: "item3", Num: 300},
	}
	for _, item := range items {
		err := r.db.Where("name = ?", item.Name).
			Assign(model.Item{Num: item.Num}).
			FirstOrCreate(&item).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *InventoryRepository) AddProcessedOrder(orderID uint) error {
	processedOrder := model.ProcessedOrder{OrderID: orderID}
	if err := r.db.Create(&processedOrder).Error; err != nil {
		return err
	}
	return nil
}

func (r *InventoryRepository) IsOrderProcessed(orderID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&model.ProcessedOrder{}).Where("order_id = ?", orderID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *InventoryRepository) ProcessOrder(orderID uint, itemname string, nums int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Model(&model.ProcessedOrder{}).Where("order_id = ?", orderID).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("order already processed")
		}
		var item model.Item
		if err := tx.Where("name = ?", itemname).First(&item).Error; err != nil {
			return err
		}
		if item.Num < nums {
			return errors.New("not enough inventory")
		}
		item.Num -= nums
		if err := tx.Save(&item).Error; err != nil {
			return err
		}
		processedOrder := model.ProcessedOrder{OrderID: orderID}
		if err := tx.Create(&processedOrder).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *InventoryRepository) PreDeduct(itemName string, nums int) error {
	result := r.db.Model(&model.Item{}).
		Where("name = ? AND num >= ?", itemName, nums).
		UpdateColumn("num", gorm.Expr("num - ?", nums))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("inventory shortage or item not found")
	}
	return nil
}

func (r *InventoryRepository) RollbackDeduct(itemName string, nums int) error {
	result := r.db.Model(&model.Item{}).
		Where("name = ?", itemName).
		UpdateColumn("num", gorm.Expr("num + ?", nums))

	return result.Error
}
