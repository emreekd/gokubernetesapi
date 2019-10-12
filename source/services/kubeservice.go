package services

import (
	"../contract/response"
	"../domain/executer"
	"../domain/repository"
)

type IKubeService interface {
	GetAllPods() response.GetAllPodsResponse
}

type kubeService struct {
	kubePodRepository  repository.IKubePodRepository
	sshCommandExecuter executer.ISshCommandExecuter
}

func InitKubeService(kpr repository.IKubePodRepository, ce executer.ISshCommandExecuter) IKubeService {
	ks := &kubeService{
		kubePodRepository:  kpr,
		sshCommandExecuter: ce,
	}
	return ks
}

func (ks *kubeService) GetAllPods() response.GetAllPodsResponse {
	var _ = ks.sshCommandExecuter.Execute("")
	var resp = new(response.GetAllPodsResponse)
	var entities = ks.kubePodRepository.GetAll()
	for _, entity := range *entities {
		newPod := &response.Pod{
			Name:     entity.Name,
			Status:   entity.Status,
			Restarts: entity.Restarts,
			Age:      entity.Age,
		}
		resp.Pods = append(resp.Pods, *newPod)
	}
	return *resp
}
