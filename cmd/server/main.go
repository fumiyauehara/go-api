package main

import (
	"fmt"
	"github.com/fumiyauehara/go-api/internal/api"
	"log"
	"net/http"
)

func main() {
	r := api.InitRouter()
	fmt.Println("Server is starting on port 3500...")
	log.Fatal(http.ListenAndServe(":3500", r))
}
