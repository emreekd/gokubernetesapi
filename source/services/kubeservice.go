package services

import (
	"../contract/response"
	"../domain/repository"
)

type IKubeService interface {
	GetAllPods() response.GetAllPodsResponse
	GetByNamespace(namespace string) response.GetAllPodsResponse
	GetDeployments(namespace string) response.GetDeploymentsResponse
	GetPortForwardCommand(podname string, namespace string, destinationPort string, localPort string) string
	GetNamespaces() response.GetNamespacesResponse
	GetNodes() response.GetNodesResponse
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

func (ks *kubeService) GetByNamespace(namespace string) response.GetAllPodsResponse {
	var resp = new(response.GetAllPodsResponse)
	var entities = ks.kubePodRepository.GetByNamespace(namespace)
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

func (ks *kubeService) GetNodes() response.GetNodesResponse {
	var resp = new(response.GetNodesResponse)
	var entities = ks.kubePodRepository.GetNodes()
	for _, entity := range *entities {
		newNode := &response.Node{
			InternalIp: entity.InternalIp,
		}
		resp.Nodes = append(resp.Nodes, *newNode)
	}
	return *resp
}

func (ks *kubeService) GetNamespaces() response.GetNamespacesResponse {
	var resp = new(response.GetNamespacesResponse)
	var entities = ks.kubePodRepository.GetNamespaces()
	for _, entity := range *entities {
		newDep := &response.Namespace{
			Name:   entity.Name,
			Age:    entity.Age,
			Status: entity.Status,
		}
		resp.Namespaces = append(resp.Namespaces, *newDep)
	}
	return *resp
}

func (ks *kubeService) GetDeployments(namespace string) response.GetDeploymentsResponse {
	var resp = new(response.GetDeploymentsResponse)
	var entities = ks.kubePodRepository.GetDeployments(namespace)
	for _, entity := range *entities {
		newDep := &response.Deployment{
			Name:          entity.Name,
			Ready:         entity.Ready,
			UpToDate:      entity.UpToDate,
			Available:     entity.Available,
			Age:           entity.Age,
			ContainerName: entity.ContainerName,
		}
		resp.Deployments = append(resp.Deployments, *newDep)
	}
	return *resp
}

func (ks *kubeService) GetPortForwardCommand(podname string, namespace string, destinationPort string, localPort string) string {
	return "kubectl port-forward " + podname + " -n " + namespace + " " + destinationPort + ":" + localPort
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
