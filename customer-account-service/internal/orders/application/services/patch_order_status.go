package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/orders/domain"
	"customer-account-service/customer-account-service/internal/orders/dtos"
	"github.com/google/uuid"
)

func PutOrderStatus(ctx context.Context, orderRepo domain.OrderRepository, statusReq dtos.StatusRequestDto, id uuid.UUID, customerId uuid.UUID) error {
	return orderRepo.Transaction(func(txRepo domain.OrderRepository) error {

		existingOrder, err := txRepo.FindById(ctx, id, customerId)

		if err != nil {
			return err
		}

		existingOrder.Status = statusReq.Status
		existingOrder.CustomerId = customerId

		return txRepo.Update(ctx, existingOrder)
	})
}
