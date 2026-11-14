package main

import (
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type WalletItem interface {
	GetBalance() (money.Money, error)
}

type Wallet struct {
	id          uuid.UUID
	ownerID     uuid.UUID
	walletItems []WalletItem
}

func (w Wallet) GetWalletBalance() (*money.Money, error) {
	var bal *money.Money
	for _, v := range w.walletItems {
		itemBal, err := v.GetBalance()
		if err != nil {
			return nil, errors.New("failed to get balance")
		}
		bal, err = bal.Add(&itemBal)
		if err != nil {
			return nil, errors.New("failed to increment balance")
		}
	}
	return bal, nil
}
