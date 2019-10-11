package persistance

import (
	"../../domain/entity"
	"../../domain/repository"
)

type kubePodRepository struct {
}

func InitKubePodRepository() repository.IKubePodRepository {
	repo := &kubePodRepository{}

	return repo
}

func (r *kubePodRepository) GetAll() *[]entity.KubePod {
	entities := &[]entity.KubePod{
		entity.KubePod{
			Name: "kubepod",
		},
	}

	return entities
}
