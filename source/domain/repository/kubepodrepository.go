package repository

import (
	"../entity"
)

type IKubePodRepository interface {
	GetAll() *[]entity.KubePod
}
