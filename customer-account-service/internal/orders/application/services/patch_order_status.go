package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/orders/domain"
	"customer-account-service/customer-account-service/internal/orders/dtos"
	"github.com/google/uuid"
)

func PatchOrderStatus(ctx context.Context, orderRepo domain.OrderRepository, statusReq dtos.StatusRequestDto, id uuid.UUID, customerId uuid.UUID) (dtos.OneOrderResponse, error) {

	var orderResponse dtos.OneOrderResponse

	if err := orderRepo.Transaction(func(txRepo domain.OrderRepository) error {

		existingOrder, err := txRepo.FindById(ctx, id, customerId)

		if err != nil {
			return err
		}

		existingOrder.Status = statusReq.Status
		existingOrder.CustomerId = customerId

		updatedOrder, err := txRepo.Patch(ctx, existingOrder)

		if err != nil {
			return err
		}

		orderResponse = dtos.OrderToDto(updatedOrder)

		return nil

	}); err != nil {
		return dtos.OneOrderResponse{}, err
	}

	return orderResponse, nil
}
