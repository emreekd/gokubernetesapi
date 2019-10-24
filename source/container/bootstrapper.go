package container

import (
	"../infrastructure/configuration"
	"../services"
	"./modules"
)

type Bootstrapper interface {
	Initialize() *services.IKubeService
}

func Initialize() services.IKubeService {

	var appSettings, _ = configuration.GetConfig("qa")
	var hostInfo = appSettings.HostInfo

	var sshCommandBuilder = modules.LoadBuilders()
	var sshClient = modules.LoadClients() // pass this client to executer
	var sshExecuter = modules.LoadExecuters(sshCommandBuilder, sshClient)
	var kubePodRepo = modules.LoadRepositories(sshExecuter, hostInfo)
	var kubeService = modules.LoadServices(kubePodRepo)

	return kubeService
}
