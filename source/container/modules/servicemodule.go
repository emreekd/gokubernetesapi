package modules

import (
	"../../domain/executer"
	"../../domain/repository"
	serviceModule "../../services"
)

type ServiceModule interface {
	LoadServices(kpr repository.IKubePodRepository) *serviceModule.IKubeService
}

func LoadServices(kpr repository.IKubePodRepository, ce executer.ISshCommandExecuter) serviceModule.IKubeService {
	var kubeService = serviceModule.InitKubeService(kpr, ce)
	return kubeService
}
