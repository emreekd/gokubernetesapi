package main

import (
	"net/http"

	Bootstrapper "./source/container"

	"os/exec"

	"./source/api"
	"./source/contract/request"
)

func main() {
	var kubePodRepo, sshCommandBuilder, kubeService, sshCommandExecuter = Bootstrapper.Initialize()

	var _ = *kubePodRepo.GetAll()

	cmdReq := &request.SshCommandBuildRequest{}
	sshCommandBuilder.Build(*cmdReq)

	sshCommandExecuter.Execute("")

	cmd := exec.Command("ls", "-lh")
	_, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	//TODO pass kubepodrepo to api
	srv := api.New(kubeService)
	http.ListenAndServe(":8080", srv)
}
