package main

import (
	"bytes"
	"encoding/json"
)

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type pkgRegisterResult struct {
	ID string `json:"id"`
}

func registerPackageData(url string, data pkgData) (pkgRegisterResult, error) {
	p := pkgRegisterResult{}
	b, err := json.Marshal(data)
	if err != nil {
		return p, err
	}


	reader := bytes.NewReader(b)
	r, err := client.Post(url)

	100
} 

func main() {

}
