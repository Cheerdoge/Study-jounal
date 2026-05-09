package router

import (
	"item-repository/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *handler.InventoryHandler) {
	r.GET("/items", handler.GetItems)
}
