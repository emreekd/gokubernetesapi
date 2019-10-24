package modules

import (
	"../../domain/executer"
	"../../domain/repository"
	"../../infrastructure/configuration"
	"../../infrastructure/persistance"
)

type RepositoryModule interface {
	LoadRepositories(ce executer.ISshCommandExecuter, hi configuration.HostInfo) *repository.IKubePodRepository
}

func LoadRepositories(ce executer.ISshCommandExecuter, hi configuration.HostInfo) repository.IKubePodRepository {
	var repo = persistance.InitKubePodRepository(ce, hi)
	return repo
}
