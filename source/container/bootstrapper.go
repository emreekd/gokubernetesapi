package container

import (
	"../domain/builder"
	"../domain/executer"
	"../domain/repository"
	"../services"
	"./modules"
)

type Bootstrapper interface {
	Initialize() (*repository.IKubePodRepository, *builder.ISshCommandBuilder, *services.IKubeService, *executer.ISshCommandExecuter)
}

func Initialize() (repository.IKubePodRepository, builder.ISshCommandBuilder, services.IKubeService, executer.ISshCommandExecuter) {
	var kubePodRepo = modules.LoadRepositories()
	var sshCommandBuilder = modules.LoadBuilders()
	var sshExecuter = modules.LoadExecuters(sshCommandBuilder)
	var kubeService = modules.LoadServices(kubePodRepo, sshExecuter)
	return kubePodRepo, sshCommandBuilder, kubeService, sshExecuter
}
