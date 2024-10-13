package main

import (
	"fmt"
	"github.com/fumiyauehara/go-api/internal/api"
	"github.com/spf13/pflag"
	"log"
	"net/http"
)

var (
	Version = "dev"
	Commit  = "none"
	BuiltBy = "unknown"
	envPath string
)

func init() {
	pflag.StringVarP(&envPath, "env-path", "e", ".env", "Path to the .env file (short)")

	pflag.Parse()
}

func main() {
	c := api.InitConfig(envPath)
	r := api.InitRouter()
	fmt.Printf("Server is starting on port %d... version: %s, commit: %s, built by: %s\n", c.Port, Version, Commit, BuiltBy)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), r))
}
