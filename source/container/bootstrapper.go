package container

import (
	"../services"
	"./modules"
)

type Bootstrapper interface {
	Initialize() *services.IKubeService
}

func Initialize() services.IKubeService {
	var sshCommandBuilder = modules.LoadBuilders()
	var sshClient = modules.LoadClients() // pass this client to executer
	var sshExecuter = modules.LoadExecuters(sshCommandBuilder, sshClient)
	var kubePodRepo = modules.LoadRepositories(sshExecuter)
	var kubeService = modules.LoadServices(kubePodRepo)

	return kubeService
}
