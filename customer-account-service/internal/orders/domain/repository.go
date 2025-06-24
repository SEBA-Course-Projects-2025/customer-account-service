package domain

import (
	"context"
	"customer-account-service/customer-account-service/internal/orders/domain/models"
	"customer-account-service/customer-account-service/internal/orders/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll(ctx context.Context, params dtos.OrderQueryParams, customerId uuid.UUID) ([]models.Order, error)
	FindById(ctx context.Context, id uuid.UUID, customerId uuid.UUID) (*models.Order, error)
	Patch(ctx context.Context, updatedOrder *models.Order) (*models.Order, error)
	Transaction(fn func(txRepo OrderRepository) error) error
	WithTx(tx *gorm.DB) OrderRepository
}
