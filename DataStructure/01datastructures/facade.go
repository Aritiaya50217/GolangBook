package main

import "fmt"

// Account struct
type Account struct {
	id          string
	accountType string
}

// Account class method create - creates account given AccountType
func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

// Account class method getById given id string
func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by id")
	return account
}

// Account class method deleteById given id string
func (account *Account) deleteById(id string) {
	fmt.Println("delete account by id")
}

// Customer struct
type Customer struct {
	name string
	id   int
}

// Customer class method create - create Customer given name
func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

// Transaction struct
type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

// Transaction class method create Transaction
func (transaction *Transaction) crate(srcAccountId, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

// BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

// method NewBranchManagerFacade
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

// BranchManagerFacade class method createCustomerAccount
func (facade *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)

	return customer, account
}

// BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.crate(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account
	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println("-------- createCustomerAccount -------- ")
	fmt.Println("name : ", customer.name)
	fmt.Println("account : ", account.accountType)

	fmt.Println("-------- createTransaction -------- ")
	var transaction = facade.createTransaction("21456", "87345", 10000)
	fmt.Println(transaction.amount)

}
