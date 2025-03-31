package main

import (
	"DataStructure/02gettingStarted/CRUDWebForms/web"
	"html/template"
	"log"
	"net/http"
)

func Home(writer http.ResponseWriter, reader *http.Request) {
	template_html := template.Must(template.ParseFiles("main.html"))
	template_html.Execute(writer, nil)

}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", web.Home)
	http.ListenAndServe(":8000", nil)
}
