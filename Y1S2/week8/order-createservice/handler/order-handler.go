package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	UserID uint   `json:"user_id"`
	Item   string `json:"item"`
}

type OrderService interface {
	CreateOrder(item string) error
}

type OrderHandler struct {
	service OrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Failed to bind request: %v", err)
		return
	}
	err := h.service.CreateOrder(req.Item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("Failed to create order: %v", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})

}
