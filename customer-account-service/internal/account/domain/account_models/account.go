package account_models

import (
	"github.com/google/uuid"
	"time"
)

type CustomerAccount struct {
	Id              uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	Email           string    `json:"email" gorm:"column:email;type:varchar(255);not null unique"`
	PasswordHash    string    `json:"password_hash" gorm:"column:password_hash;type:varchar(255);not null"`
	Phone           string    `json:"phone" gorm:"column:phone;type:varchar(255);not null unique"`
	Name            string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	ShippingAddress string    `json:"shipping_address" gorm:"column:shipping_address;type:varchar(255);not null"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp"`
}
