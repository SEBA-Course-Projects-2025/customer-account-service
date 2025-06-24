package services

import (
	"context"
	"customer-account-service/customer-account-service/internal/account/domain"
	"customer-account-service/customer-account-service/internal/account/dtos"
	"customer-account-service/customer-account-service/internal/shared/utils/jwt_helper"
	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(ctx context.Context, accountRepo domain.AccountRepository, loginReq dtos.LoginRequestDto) (string, error) {

	account, err := accountRepo.FindByEmail(ctx, loginReq.Email)

	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(loginReq.Password)); err != nil {
		return "", err
	}

	return jwt_helper.GenerateCustomerJwt(account.Id)

}
