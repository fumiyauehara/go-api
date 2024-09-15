package handler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "Hello World"); err != nil {
		panic(err)
	}
}
