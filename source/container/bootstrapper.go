package container

import (
	"../domain/builder"
	"../domain/repository"
	"../services"
	"./modules"
)

type Bootstrapper interface {
	Initialize() (*repository.IKubePodRepository, *builder.ISshCommandBuilder, *services.IKubeService)
}

func Initialize() (repository.IKubePodRepository, builder.ISshCommandBuilder, services.IKubeService) {
	var sshCommandBuilder = modules.LoadBuilders()
	var sshExecuter = modules.LoadExecuters(sshCommandBuilder)
	var kubePodRepo = modules.LoadRepositories(sshExecuter)
	var kubeService = modules.LoadServices(kubePodRepo)
	return kubePodRepo, sshCommandBuilder, kubeService
}
