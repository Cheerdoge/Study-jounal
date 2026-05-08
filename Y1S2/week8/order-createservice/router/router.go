package router

import (
	"order-payment-kafka/order-createservice/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, orderhandler *handler.OrderHandler) {
	r.POST("/orders", orderhandler.CreateOrder)
}
