package modules

import (
	"../../domain/executer"
	"../../domain/repository"
	"../../infrastructure/persistance"
)

type RepositoryModule interface {
	LoadRepositories(ce executer.ISshCommandExecuter) *repository.IKubePodRepository
}

func LoadRepositories(ce executer.ISshCommandExecuter) repository.IKubePodRepository {
	var repo = persistance.InitKubePodRepository(ce)
	return repo
}
