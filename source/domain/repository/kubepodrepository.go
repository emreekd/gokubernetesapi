package repository

import (
	"../entity"
)

type IKubePodRepository interface {
	GetAll() *[]entity.KubePod
	GetByNamespace(namespace string) *[]entity.KubePod
}
