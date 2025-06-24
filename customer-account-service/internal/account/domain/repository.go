package domain

import (
	"context"
	"customer-account-service/customer-account-service/internal/account/domain/account_models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindById(ctx context.Context, customerId uuid.UUID) (*account_models.CustomerAccount, error)
	Update(ctx context.Context, updatedAccount *account_models.CustomerAccount) error
	Patch(ctx context.Context, modifiedAccount *account_models.CustomerAccount) (*account_models.CustomerAccount, error)
	FindByEmail(ctx context.Context, email string) (*account_models.CustomerAccount, error)
	Transaction(fn func(txRepo AccountRepository) error) error
	WithTx(tx *gorm.DB) AccountRepository
}
