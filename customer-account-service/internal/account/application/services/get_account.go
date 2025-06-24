package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/account/domain"
	"customer-account-service/customer-account-service/internal/account/dtos"
	"github.com/google/uuid"
)

func GetAccount(ctx context.Context, repo domain.AccountRepository, customerId uuid.UUID) (dtos.AccountResponse, error) {

	account, err := repo.FindById(ctx, customerId)

	if err != nil {
		return dtos.AccountResponse{}, nil
	}

	return dtos.AccountToDto(account), nil
}
