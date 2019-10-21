package repository

import (
	"../entity"
)

type IKubePodRepository interface {
	GetAll() *[]entity.KubePod
	GetByNamespace(namespace string) *[]entity.KubePod
	GetDeployments(namespace string) *[]entity.Deployment
	GetNamespaces() *[]entity.Namespace
	GetNodes() *[]entity.Node
}
