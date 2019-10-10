package repository

import (
	"../entity"
)

type KubePodRepository interface {
	GetAll() *entity.KubePod
}

var kubePodRepository KubePodRepository

func InitRepository(r KubePodRepository) {
	kubePodRepository = r
}

func GetRepository() KubePodRepository {
	return kubePodRepository
}
