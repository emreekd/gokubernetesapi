package main

import (
	"net/http"

	"./source/api"
)

func main() {
	srv := api.New()
	http.ListenAndServe(":8080", srv)
}
