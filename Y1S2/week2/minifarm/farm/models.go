package farm

import (
	"context"
	"sync"
)

type State string
type Action string

const (
	StateWaitWater     State = "等待浇水"
	StateWaitWeed      State = "等待除草虫"
	StateWaitFertilize State = "等待施肥"
	StateWaitHarvest   State = "等待收获"
	StateHarvested     State = "已收获"
	StateDead          State = "已枯死"

	ActionWater     Action = "浇水"
	ActionWeed      Action = "除草虫"
	ActionFertilize Action = "施肥"
	ActionHarvest   Action = "收获"
)

// CheckValidAction 检查当前状态是否允许该动作
func CheckValidAction(state State, action Action) bool {
	switch state {
	case StateWaitWater:
		return action == ActionWater
	case StateWaitWeed:
		return action == ActionWeed
	case StateWaitFertilize:
		return action == ActionFertilize
	case StateWaitHarvest:
		return action == ActionHarvest
	default:
		return false
	}
}

// Plant 表示一棵蔬菜
type Plant struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	State State  `json:"state"`

	// 并发控制与通信
	ActionCh  chan Action        `json:"-"`
	Ctx       context.Context    `json:"-"`
	CancelCtx context.CancelFunc `json:"-"`
	Mutex     sync.RWMutex       `json:"-"` // 保护 State 的并发读写
}

// GetState 线程安全地获取当前状态
func (p *Plant) GetState() State {
	p.Mutex.RLock()
	defer p.Mutex.RUnlock()
	return p.State
}

// SetState 线程安全地设置当前状态
func (p *Plant) SetState(s State) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	p.State = s
}
