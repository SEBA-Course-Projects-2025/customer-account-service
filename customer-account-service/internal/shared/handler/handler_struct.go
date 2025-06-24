package handlers

import (
	accountDomain "customer-account-service/customer-account-service/internal/account/domain"
	orderDomain "customer-account-service/customer-account-service/internal/orders/domain"
	"gorm.io/gorm"
)

type Handler struct {
	AccountRepo accountDomain.AccountRepository

	OrderRepo orderDomain.OrderRepository

	Db *gorm.DB
}
