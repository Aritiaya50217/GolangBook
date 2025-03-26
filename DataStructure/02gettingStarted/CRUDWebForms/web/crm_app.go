package web

import (
	"html/template"
	"log"
	"net/http"
)

var template_html = template.Must(template.ParseGlob("./templates/*"))

// Home - execute Template
func Home(writer http.ResponseWriter, reader *http.Request) {
	customers := GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", "customers")
}
// test