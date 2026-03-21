package farm

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PlantReq 种植请求
type PlantReq struct {
	Type string `json:"type" binding:"required"`
}

// ActionReq 互动请求
type ActionReq struct {
	ID     string `json:"id" binding:"required"`
	Action Action `json:"action" binding:"required"`
}

// FarmHandlers Gin 的处理函数结构
type FarmHandlers struct {
	fm *FarmManager
}

func NewFarmHandlers(fm *FarmManager) *FarmHandlers {
	return &FarmHandlers{fm: fm}
}

// RegisterRoutes 注册路由
func (h *FarmHandlers) RegisterRoutes(r *gin.Engine) {
	farmGroup := r.Group("/api/farm")
	{
		farmGroup.POST("/plant", h.HandlePlant)
		farmGroup.POST("/action", h.HandleAction)
		farmGroup.GET("/status", h.HandleStatus)
	}
}

// HandlePlant 种植一棵蔬菜
func (h *FarmHandlers) HandlePlant(c *gin.Context) {
	var req PlantReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误，需要蔬菜名称 type"})
		return
	}

	plant := h.fm.PlantVegetable(req.Type)
	c.JSON(http.StatusOK, gin.H{
		"id":      plant.ID,
		"msg":     fmt.Sprintf("成功种下 [%s]，请在 30 秒内进行[%s]", plant.Type, ActionWater),
		"plant":   plant.Type,
		"current": string(StateWaitWater),
	})
}

// HandleAction 进行互动（浇水、除草虫、施肥、收获）
func (h *FarmHandlers) HandleAction(c *gin.Context) {
	var req ActionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误，需要 id 和 action"})
		return
	}

	plant, err := h.fm.GetPlant(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	currentState := plant.GetState()

	// 1. 如果已经枯死或已收获，直接报错
	if currentState == StateDead {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该蔬菜已经枯死，无法操作"})
		return
	}
	if currentState == StateHarvested {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该蔬菜已经收获，无法操作"})
		return
	}

	// 2. 严格校验状态是否允许该动作
	if !CheckValidAction(currentState, req.Action) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("动作不合法！当前处于 [%s] 状态，不能执行 [%s] 操作", currentState, req.Action),
		})
		return
	}

	// 3. 将动作推入通道
	// 使用非阻塞推送（select-default 模式防止 channel 满导致协程阻塞在 API 层）
	// 因为前面我们 make(chan Action, 1) 设置了容量，所以一般不会满
	select {
	case plant.ActionCh <- req.Action:
		c.JSON(http.StatusOK, gin.H{
			"msg":    fmt.Sprintf("对蔬菜 [%s] 执行了 [%s] 动作成功！", plant.Type, req.Action),
			"action": req.Action,
		})
	default:
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "蔬菜正忙或通道已满，请稍后再试"})
	}
}

// HandleStatus 查看农场中所有蔬菜状态
func (h *FarmHandlers) HandleStatus(c *gin.Context) {
	plants := h.fm.GetAllPlants()
	c.JSON(http.StatusOK, gin.H{
		"total":  len(plants),
		"plants": plants,
	})
}
