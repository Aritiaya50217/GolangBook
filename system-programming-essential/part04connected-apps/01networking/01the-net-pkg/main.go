package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://pokeapi.co/api/v2/pokemon/ditto"
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	fmt.Println("http status code : ", resp.StatusCode)
	fmt.Println("response : ", string(body))

}
