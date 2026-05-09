package handler

import (
	"context"
	inventory "item-repository/pb"
	"item-repository/repository"
)

type InventoryGrpcHandler struct {
	inventory.UnimplementedInventoryServiceServer
	repo *repository.InventoryRepository
}

func NewInventoryGrpcHandler(repo *repository.InventoryRepository) *InventoryGrpcHandler {
	return &InventoryGrpcHandler{repo: repo}
}
func (h *InventoryGrpcHandler) PreDeduct(ctx context.Context, req *inventory.PreDeductRequest) (*inventory.PreDeductResponse, error) {
	err := h.repo.PreDeduct(req.ProductId, int(req.Quantity))
	if err != nil {
		return &inventory.PreDeductResponse{Success: false, Message: err.Error()}, nil
	}
	return &inventory.PreDeductResponse{Success: true, Message: "success"}, nil
}
func (h *InventoryGrpcHandler) RollbackDeduct(ctx context.Context, req *inventory.RollbackDeductRequest) (*inventory.RollbackDeductResponse, error) {
	err := h.repo.RollbackDeduct(req.ProductId, int(req.Quantity))
	if err != nil {
		return &inventory.RollbackDeductResponse{Success: false, Message: err.Error()}, nil
	}
	return &inventory.RollbackDeductResponse{Success: true, Message: "success"}, nil
}
