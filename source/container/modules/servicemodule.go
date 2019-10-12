package modules

import (
	"../../domain/repository"
	serviceModule "../../services"
)

type ServiceModule interface {
	LoadServices(kpr repository.IKubePodRepository) *serviceModule.IKubeService
}

func LoadServices(kpr repository.IKubePodRepository) serviceModule.IKubeService {
	var kubeService = serviceModule.InitKubeService(kpr)
	return kubeService
}
