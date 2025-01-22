package pkgquery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func fetchPackageData(url string) ([]pkgData, error) {
	var packages []pkgData
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	if r.Header.Get("Content-Type") != "application/json" {
		return packages, nil
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return packages, err
	}

	err = json.Unmarshal(data, &packages)

	return packages, err
}

func startTestPackageServer() *httptest.Server {
	pkgData := `[
	   {"name": "package1", "version": "1.1"},
	   {"name": "package2", "version": "1.0"}
	   ]`

	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, pkgData)
			},
		),
	)
	return ts
}
