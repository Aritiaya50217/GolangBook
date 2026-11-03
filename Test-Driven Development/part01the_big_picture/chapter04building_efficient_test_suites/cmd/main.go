package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"test-driven-development/part01the_big_picture/chapter04building_efficient_test_suites/db"
	"test-driven-development/part01the_big_picture/chapter04building_efficient_test_suites/handlers"
)

var booksFile []byte
var usersFile []byte

func main() {
	port, ok := os.LookupEnv("BOOKSWAP_PORT")
	if !ok {
		log.Fatal("$BOOKSWAP_PORT not found")
	}

	books, users := importInitial()
	ps := db.NewPostingService()
	b := db.NewBookService(books, ps)
	u := db.NewUserService(users, b)
	h := handlers.NewHandler(b, u)

	router := handlers.ConfigureServer(h)
	log.Printf("Listening on :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}

func importInitial() (books []db.Book, users []db.User) {
	if err := json.Unmarshal(booksFile, &books); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(usersFile, users); err != nil {
		log.Fatal(err)
	}

	return books, users

}
