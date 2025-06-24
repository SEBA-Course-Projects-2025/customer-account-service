package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/orders/domain"
	"customer-account-service/customer-account-service/internal/orders/dtos"
	"github.com/google/uuid"
)

func GetOrderById(ctx context.Context, orderRepo domain.OrderRepository, id uuid.UUID, customerId uuid.UUID) (dtos.OneOrderResponse, error) {

	order, err := orderRepo.FindById(ctx, id, customerId)

	if err != nil {
		return dtos.OneOrderResponse{}, err
	}

	return dtos.OrderToDto(order), nil
}
