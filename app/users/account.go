package users

import (
	"github.com/1layar/clibank/app/wallets"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name        string          `gorm:"type:varchar(100)"`
	PhoneNumber string          `gorm:"type:varchar(12);uniqueIndex;not null" json:"phone_number"`
	WalletID    uint            `json:"-"`
	Wallet      *wallets.Wallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"wallet"`
}
