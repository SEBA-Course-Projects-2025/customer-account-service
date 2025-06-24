package dtos

import (
	"customer-account-service/customer-account-service/internal/account/domain/account_models"
	"github.com/google/uuid"
	"time"
)

type AccountResponse struct {
	Id              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	ShippingAddress string    `json:"address"`
}

func AccountToDto(account *account_models.CustomerAccount) AccountResponse {
	return AccountResponse{
		Id:              account.Id,
		Email:           account.Email,
		Name:            account.Name,
		Phone:           account.Phone,
		ShippingAddress: account.ShippingAddress,
	}
}

type PutRequestDto struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	ShippingAddress string `json:"address"`
}

func UpdateAccountWithDto(accountReq PutRequestDto, existingAccount *account_models.CustomerAccount) account_models.CustomerAccount {
	return account_models.CustomerAccount{
		Id:              existingAccount.Id,
		Email:           accountReq.Email,
		PasswordHash:    existingAccount.PasswordHash,
		Name:            accountReq.Name,
		Phone:           accountReq.Phone,
		ShippingAddress: accountReq.ShippingAddress,
		CreatedAt:       existingAccount.CreatedAt,
		UpdatedAt:       time.Now(),
	}
}

type AccountPatchRequest struct {
	Email           *string `json:"email"`
	Name            *string `json:"name"`
	Phone           *string `json:"phone"`
	ShippingAddress *string `json:"address"`
}

func PatchDtoToAccount(existingAccount *account_models.CustomerAccount, accountReq AccountPatchRequest) *account_models.CustomerAccount {

	if accountReq.Email != nil {
		existingAccount.Email = *accountReq.Email
	}

	if accountReq.Name != nil {
		existingAccount.Name = *accountReq.Name
	}

	if accountReq.Phone != nil {
		existingAccount.Phone = *accountReq.Phone
	}

	if accountReq.ShippingAddress != nil {
		existingAccount.ShippingAddress = *accountReq.ShippingAddress
	}

	existingAccount.UpdatedAt = time.Now()

	return existingAccount
}

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
