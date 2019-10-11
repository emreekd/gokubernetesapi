package modules

import (
	"../../domain/repository"
	"../../infrastructure/persistance"
)

type RepositoryModule interface {
	LoadRepositories() *repository.IKubePodRepository
}

func LoadRepositories() repository.IKubePodRepository {
	var repo = persistance.InitKubePodRepository()
	return repo
}
