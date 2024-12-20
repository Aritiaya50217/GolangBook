package main

import "fmt"

// DataTransferObjectFactory struct
type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

// DataTransferObjectFactory class method getDataTransferObject
func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO of dtoType : " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = CustomerForFlyweigh{id: "1"}
		case "employee":
			factory.pool[dtoType] = EmployeeForFlyweigh{id: "2"}
		case "manager":
			factory.pool[dtoType] = ManagerForFlyweigh{id: "3"}
		case "address":
			factory.pool[dtoType] = AddressForFlyweigh{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

// DataTransferObject interface
type DataTransferObject interface {
	getId() string
}

type CustomerForFlyweigh struct {
	id   string //sequence generator
	name string
	ssn  string
}

// Customer class method getId
func (customer CustomerForFlyweigh) getId() string {
	return customer.id
}

type EmployeeForFlyweigh struct {
	id   string
	name string
}

// Employee class method getId
func (employee EmployeeForFlyweigh) getId() string {
	return employee.id
}

// Manager struct
type ManagerForFlyweigh struct {
	id   string
	name string
	dept string
}

// Manager class method getId
func (manager ManagerForFlyweigh) getId() string {
	return manager.id
}

// Address
type AddressForFlyweigh struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func (address AddressForFlyweigh) getId() string {
	return address.id
}

func main() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}
	var customer DataTransferObject = factory.getDataTransferObject("customer")
	fmt.Println("Customer id : ", customer.getId())
	var employee DataTransferObject = factory.getDataTransferObject("employee")
	fmt.Println("Employee id : ", employee.getId())
	var manager DataTransferObject = factory.getDataTransferObject("manager")
	fmt.Println("Manager id : ", manager.getId())

}
