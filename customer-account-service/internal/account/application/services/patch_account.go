package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/account/domain"
	"customer-account-service/customer-account-service/internal/account/dtos"
	"github.com/google/uuid"
)

func PatchAccount(ctx context.Context, accountRepo domain.AccountRepository, accountReq dtos.AccountPatchRequest, customerId uuid.UUID) (dtos.AccountResponse, error) {

	var accountResponse dtos.AccountResponse

	if err := accountRepo.Transaction(func(txRepo domain.AccountRepository) error {

		existingAccount, err := txRepo.FindById(ctx, customerId)

		if err != nil {
			return err
		}

		existingAccount = dtos.PatchDtoToAccount(existingAccount, accountReq)

		existingAccount.Id = customerId

		existingAccount, err = txRepo.Patch(ctx, existingAccount)

		if err != nil {
			return err
		}

		accountResponse = dtos.AccountToDto(existingAccount)

		return nil

	}); err != nil {
		return dtos.AccountResponse{}, err
	}

	return accountResponse, nil
}
