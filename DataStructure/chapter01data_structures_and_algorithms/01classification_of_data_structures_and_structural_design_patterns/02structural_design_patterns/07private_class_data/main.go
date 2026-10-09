package main

import (
	"encoding/json"
	"fmt"
)

// AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

type AccountPrivate struct {
	details      *AccountDetails
	CustomerName string
}

func (account *AccountPrivate) setDetails(id, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

func (account *AccountPrivate) getId() string {
	return account.details.id
}

func (account *AccountPrivate) getAccountType() string {
	return account.details.accountType
}

func main() {
	var account *AccountPrivate = &AccountPrivate{CustomerName: "John"}
	account.setDetails("1234", "current")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))
	fmt.Println("Account Id : ", account.getId())
	fmt.Println("Account Type ", account.getAccountType())
}
