package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *InventoryHandler) GetItems(c *gin.Context) {
	items, err := h.Service.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
