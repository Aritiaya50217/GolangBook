package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var template_html = template.Must(template.ParseGlob("./templates/*"))

// Home - execute Template
func Home(writer http.ResponseWriter, reader *http.Request) {
	customers := GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	template_html.ExecuteTemplate(writer, "Create", nil)
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	InsertCustomer(customer)
	customers := GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Alter(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	UpdateCustomer(customer)
	customers := GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer := GetCustomerById(customerId)
	template_html.ExecuteTemplate(writer, "Update", customer)

}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer := GetCustomerById(customerId)
	DeleteCustomer(customer)
	customers := GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func View(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	customerIdStr := request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer := GetCustomerById(customerId)
	fmt.Println(customer)
	customers := []Customer{customer}
	template_html.ExecuteTemplate(writer, "View", customers)
}

func main() {
	log.Println("Server started on: http://localhost:8000")

	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}
