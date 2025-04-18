package main

import (
	"database/sql"
	"fmt"
)

// Customer class
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver,
		databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	database := GetConnection()
	insert, err := database.Prepare("INSERT INTO CUSTOMER(CustomerName , SSN) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}

// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	database := GetConnection()
	rows, err := database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if err != nil {
		panic(err.Error())
	}
	customer := Customer{}
	customers := []Customer{}
	for rows.Next() {
		customerId := 0
		customerName := ""
		ssn := ""
		err := rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}

// UpdateCustomer method with parameter customer
func UpdateCustomer(customer Customer) {
	database := GetConnection()
	if _, err := database.Prepare("UPDATE CUSTOMER SET CustomerName=?,SSN=? WHERE CustomerId=?"); err != nil {
		panic(err.Error())
	}
	var update *sql.Stmt
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	defer database.Close()
}

// DeleteCustomer
func DeleteCustomer(customer Customer) {
	database := GetConnection()
	if _, err := database.Prepare("DELETE FROM Customer WHERE Customerid=?"); err != nil {
		panic(err.Error())
	}
	var delete *sql.Stmt
	delete.Exec(customer.CustomerId)
	defer database.Close()
}

func main() {
	// get
	customers := GetCustomers()
	fmt.Println("Customers : ", customers)

	// insert
	fmt.Println("Before Insert : ", customers)
	customer := Customer{}
	customer.CustomerName = "Arnie Smith"
	customer.SSN = "2345678"
	InsertCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Insert : ", customers)

	// update
	customer.CustomerName = "George Thompson"
	customer.SSN = "23456234"
	customer.CustomerId = 5
	UpdateCustomer(customer)

	customers = GetCustomers()
	fmt.Println("After Update : ", customers)

	// delete
	fmt.Println("Before Delete", customers)
	customer.CustomerName = "George Thompson"
	customer.SSN = "23233432"
	customer.CustomerId = 5
	DeleteCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Delete", customers)

}
