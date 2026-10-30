package thetestmainapproach

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup statements
	setup()

	// run the tests
	e := m.Run()

	// cleanup statements
	teardown()

	// report the exit code
	os.Exit(e)

}

func setup() {
	log.Println("Setting up.")
}

func teardown() {
	log.Println("Tearing down. ")
}
