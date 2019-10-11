package main

import (
	"fmt"
	"net/http"

	Bootstrapper "./source/container"

	"./source/api"
)

func main() {
	var kubePodRepo = Bootstrapper.Initialize()

	var entities = *kubePodRepo.GetAll()
	fmt.Printf(entities[0].Name)

	//TODO pass kubepodrepo to api
	srv := api.New()
	http.ListenAndServe(":8080", srv)
}
