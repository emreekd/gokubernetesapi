package main

import (
	"net/http"

	Bootstrapper "./source/container"

	"./source/api"
)

func main() {
	var kubeService = Bootstrapper.Initialize()

	srv := api.New(kubeService)
	http.ListenAndServe(":8081", srv)
}
