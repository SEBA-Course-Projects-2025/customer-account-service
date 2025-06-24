package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/account/domain"
	"customer-account-service/customer-account-service/internal/account/dtos"
	"github.com/google/uuid"
)

func PutAccount(ctx context.Context, accountRepo domain.AccountRepository, accountReq dtos.PutRequestDto, customerId uuid.UUID) error {

	return accountRepo.Transaction(func(txRepo domain.AccountRepository) error {
		existingAccount, err := txRepo.FindById(ctx, customerId)

		if err != nil {
			return err
		}

		*existingAccount = dtos.UpdateAccountWithDto(accountReq, existingAccount)

		return txRepo.Update(ctx, existingAccount)
	})

}
