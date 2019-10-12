package services

import (
	"../contract/response"
	"../domain/repository"
)

type IKubeService interface {
	GetAllPods() response.GetAllPodsResponse
}

type kubeService struct {
	kubePodRepository repository.IKubePodRepository
}

func InitKubeService(kpr repository.IKubePodRepository) IKubeService {
	ks := &kubeService{
		kubePodRepository: kpr,
	}
	return ks
}

func (ks *kubeService) GetAllPods() response.GetAllPodsResponse {
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
