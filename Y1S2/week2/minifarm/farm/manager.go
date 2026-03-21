package farm

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// FarmManager 管理农场所有的植物
type FarmManager struct {
	Plants sync.Map
	Ctx    context.Context // 全局 context，可用于整体农场关闭（扩展用）
}

// NewFarmManager 初始化农场管理器
func NewFarmManager() *FarmManager {
	return &FarmManager{
		Ctx: context.Background(),
	}
}

// PlantVegetable 种植新的蔬菜
func (fm *FarmManager) PlantVegetable(vegType string) *Plant {
	id := uuid.New().String()

	ctx, cancel := context.WithCancel(fm.Ctx)

	plant := &Plant{
		ID:        id,
		Type:      vegType,
		State:     StateWaitWater,
		ActionCh:  make(chan Action, 1), // 无缓冲或容量为1，此处用1防止阻塞过久
		Ctx:       ctx,
		CancelCtx: cancel,
	}

	fm.Plants.Store(id, plant)

	// 启动这棵蔬菜的生命周期 Goroutine
	go runLifecycle(plant)

	return plant
}

// GetPlant 通过 ID 获取植物
func (fm *FarmManager) GetPlant(id string) (*Plant, error) {
	v, ok := fm.Plants.Load(id)
	if !ok {
		return nil, fmt.Errorf("找不到对应的蔬菜，ID: %s", id)
	}
	return v.(*Plant), nil
}

// GetAllPlants 获取农场里所有植物的快照
func (fm *FarmManager) GetAllPlants() []*Plant {
	var list []*Plant
	fm.Plants.Range(func(key, value interface{}) bool {
		p := value.(*Plant)

		// 为了返回给 API 安全的当前状态快照，我们可以深拷贝一下关键信息
		snapshot := &Plant{
			ID:    p.ID,
			Type:  p.Type,
			State: p.GetState(),
		}
		list = append(list, snapshot)
		return true // 遍历下一个
	})
	return list
}
