package repository

import (
	"context"
	"customer-account-service/customer-account-service/internal/account/domain"
	"customer-account-service/customer-account-service/internal/account/domain/account_models"
	"customer-account-service/customer-account-service/internal/shared/utils/error_handler"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type GormAccountRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GormAccountRepository {
	return &GormAccountRepository{db: db}
}

func (gar *GormAccountRepository) FindById(ctx context.Context, customerId uuid.UUID) (*account_models.CustomerAccount, error) {

	var account account_models.CustomerAccount

	if err := gar.db.WithContext(ctx).First(&account, "id = ?", customerId).Error; err != nil {
		return nil, error_handler.ErrorHandler(err, "Error getting customer's account data")
	}

	return &account, nil
}

func (gar *GormAccountRepository) Update(ctx context.Context, updatedAccount *account_models.CustomerAccount) error {

	if err := gar.db.WithContext(ctx).Save(updatedAccount).Error; err != nil {
		return error_handler.ErrorHandler(err, "Error updating account")
	}

	return nil

}

func (gar *GormAccountRepository) Patch(ctx context.Context, modifiedAccount *account_models.CustomerAccount) (*account_models.CustomerAccount, error) {

	if err := gar.db.WithContext(ctx).Save(modifiedAccount).Error; err != nil {
		return nil, error_handler.ErrorHandler(err, "Error modifying account")
	}

	return modifiedAccount, nil

}

func (gar *GormAccountRepository) FindByEmail(ctx context.Context, email string) (*account_models.CustomerAccount, error) {

	var account account_models.CustomerAccount

	if err := gar.db.WithContext(ctx).First(&account, "email = ?", email).Error; err != nil {
		return nil, error_handler.ErrorHandler(err, "Error getting customer's account data")
	}

	return &account, nil

}

func (gar *GormAccountRepository) WithTx(tx *gorm.DB) domain.AccountRepository {
	return &GormAccountRepository{
		db: tx,
	}
}

func (gar *GormAccountRepository) Transaction(fn func(txRepo domain.AccountRepository) error) error {
	tx := gar.db.Begin()
	if tx.Error != nil {
		log.Printf("Transaction begin error: %v", tx.Error)
		return tx.Error
	}

	repo := gar.WithTx(tx)

	if err := fn(repo); err != nil {
		log.Printf("Transaction function error: %v", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Transaction commit error: %v", err)
		return err
	}

	return nil

}
