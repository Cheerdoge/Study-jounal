package farm

import (
	"context"
	"fmt"
	"time"
)

const stageTimeout = 30 * time.Second

// runLifecycle 负责这棵蔬菜生命周期的 Goroutine
func runLifecycle(p *Plant) {
	// 任何时刻 ctx 取消（例如农场倒闭，或者被外部意外杀掉）都会退出
	defer p.CancelCtx()

	// 阶段 1：等待浇水 (WaitWater)
	// 如果超时没接收到，或者接收到其它（在API层已经拦截），则进入枯死状态。
	if !waitForAction(p, StateWaitWater, ActionWater) {
		p.SetState(StateDead)
		fmt.Printf("[蔬菜 %s(%s)] 超过30s未 %s, 已枯死\n", p.Type, p.ID, ActionWater)
		return
	}
	p.SetState(StateWaitWeed)
	fmt.Printf("[蔬菜 %s(%s)] 已 %s，进入 %s 阶段\n", p.Type, p.ID, ActionWater, StateWaitWeed)

	// 阶段 2：等待除草虫 (WaitWeed)
	if !waitForAction(p, StateWaitWeed, ActionWeed) {
		p.SetState(StateDead)
		fmt.Printf("[蔬菜 %s(%s)] 超过30s未 %s, 已枯死\n", p.Type, p.ID, ActionWeed)
		return
	}
	p.SetState(StateWaitFertilize)
	fmt.Printf("[蔬菜 %s(%s)] 已 %s，进入 %s 阶段\n", p.Type, p.ID, ActionWeed, StateWaitFertilize)

	// 阶段 3：等待施肥 (WaitFertilize)
	if !waitForAction(p, StateWaitFertilize, ActionFertilize) {
		p.SetState(StateDead)
		fmt.Printf("[蔬菜 %s(%s)] 超过30s未 %s, 已枯死\n", p.Type, p.ID, ActionFertilize)
		return
	}
	p.SetState(StateWaitHarvest)
	fmt.Printf("[蔬菜 %s(%s)] 已 %s，进入 %s 阶段\n", p.Type, p.ID, ActionFertilize, StateWaitHarvest)

	// 阶段 4：等待收获 (WaitHarvest)
	if !waitForAction(p, StateWaitHarvest, ActionHarvest) {
		p.SetState(StateDead)
		fmt.Printf("[蔬菜 %s(%s)] 成熟后超过30s未 %s, 烂在地里(枯死)\n", p.Type, p.ID, ActionHarvest)
		return
	}
	p.SetState(StateHarvested)
	fmt.Printf("[蔬菜 %s(%s)] 恭喜！已成功 %s\n", p.Type, p.ID, ActionHarvest)
}

// waitForAction 创建一个 30s 超时的 Context，监听 p.ActionCh，直到收到期望的 action 或超时。
// 返回 bool: 成功执行返回 true，超时枯死返回 false。
func waitForAction(p *Plant, expectState State, expectAction Action) bool {
	// 为当前阶段创建一个有超时时间的 context
	ctx, cancel := context.WithTimeout(p.Ctx, stageTimeout)
	defer cancel() // 当正常收到信号或者其他情况退出时，释放这个 Timer

	select {
	case <-ctx.Done(): // 超过 30s 触发超时
		return false
	case <-p.Ctx.Done(): // 农场全局关闭或蔬菜被销毁
		return false
	case act := <-p.ActionCh:
		if act == expectAction {
			// 收到了正确的动作
			return true
		}
		// 如果这里收到了不合法的动作（理论上 API 已经拦截了，属于兜底防卫）
		return false
	}
}
