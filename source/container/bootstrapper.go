package container

import (
	"../domain/repository"
	"./modules"
)

type Bootstrapper interface {
	Initialize() *repository.IKubePodRepository
}

func Initialize() repository.IKubePodRepository {
	var kubePodRepo = modules.LoadRepositories()
	return kubePodRepo
}
