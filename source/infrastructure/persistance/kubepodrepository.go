package persistance

import (
	"../../domain/entity"
)

type KubePodRepositoryImpl struct {
}

func (r *KubePodRepositoryImpl) GetAll() []*entity.KubePod {
	entities := []entity.KubePod{
		entity{
			Name: "kubepod",
		}
	}

	return &entities
}
