package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/orders/domain"
	"customer-account-service/customer-account-service/internal/orders/dtos"
	"github.com/google/uuid"
)

func GetOrders(ctx context.Context, orderRepo domain.OrderRepository, params dtos.OrderQueryParams, customerId uuid.UUID) ([]dtos.GetOrdersResponse, error) {

	orders, err := orderRepo.FindAll(ctx, params, customerId)

	if err != nil {
		return nil, err
	}

	return dtos.OrdersToDto(orders), nil

}
