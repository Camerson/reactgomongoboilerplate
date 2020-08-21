package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reactgomongo/router"
)

func main() {
	os.Setenv("TEST_ENV_VARIABLE", "An example of adding an environment variable")

	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}