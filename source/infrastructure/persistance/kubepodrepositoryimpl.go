package persistance

import (
	"strings"

	"../../domain/entity"
	"../../domain/executer"
	"../../domain/repository"
)

type kubePodRepository struct {
	sshCommandExecuter executer.ISshCommandExecuter
}

func InitKubePodRepository(ce executer.ISshCommandExecuter) repository.IKubePodRepository {
	repo := &kubePodRepository{
		sshCommandExecuter: ce,
	}

	return repo
}

func (r *kubePodRepository) GetPodInfo(podName string, namespace string) string {
	var stringResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420", "kubectl", "get pod "+podName+"  -n "+namespace+" -o json")
	return string(stringResult)
}

func (r *kubePodRepository) GetNamespaces() *[]entity.Namespace {
	entities := make([]entity.Namespace, 0)
	var stringResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420", "kubectl", "get namespaces")
	namespaceInfos := strings.Split(string(stringResult), "\n")
	for _, namespaceInfo := range namespaceInfos[1:] {
		namespaceFielads := strings.Fields(namespaceInfo)
		if namespaceFielads != nil && len(namespaceFielads) >= 2 {
			newEntitiy := entity.Namespace{
				Name:   namespaceFielads[0],
				Status: namespaceFielads[1],
				Age:    namespaceFielads[2],
			}
			entities = append(entities, newEntitiy)
		}
	}
	return &entities
}

func (r *kubePodRepository) GetDeployments(namespace string) *[]entity.Deployment {
	entities := make([]entity.Deployment, 0)
	var stringResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420", "kubectl", "get deployments -n "+namespace)
	deploymentInfos := strings.Split(string(stringResult), "\n")
	for _, deploymentInfo := range deploymentInfos[1:] {
		deploymentFields := strings.Fields(deploymentInfo)
		if deploymentFields != nil && len(deploymentFields) >= 4 {
			var containerNameResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420", "kubectl",
				"get deployment "+deploymentFields[0]+" -n "+namespace+" -o=jsonpath='{..containers[0].name}'")
			containerName := string(containerNameResult)
			newEntitiy := entity.Deployment{
				Name:          deploymentFields[0],
				Ready:         deploymentFields[1],
				UpToDate:      deploymentFields[2],
				Available:     deploymentFields[3],
				Age:           deploymentFields[4],
				ContainerName: containerName,
			}
			entities = append(entities, newEntitiy)
		}
	}
	return &entities
}

func (r *kubePodRepository) UpdateImageForDeployment(deploymentName string, containerName string, newImage string, namespace string) bool {
	var stringResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420",
		"kubectl", "set image deployment/"+deploymentName+" "+containerName+"="+newImage+" -n "+namespace)
	result := string(stringResult)
	if result == "" {
		return true
	}
	return false
}

func (r *kubePodRepository) GetByNamespace(namespace string) *[]entity.KubePod {
	entities := make([]entity.KubePod, 0)
	var stringResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420", "kubectl", "get pods -n "+namespace)
	podInfos := strings.Split(string(stringResult), "\n")
	for _, podInfo := range podInfos[1:] {
		podFields := strings.Fields(podInfo)
		if podFields != nil && len(podFields) > 2 {
			newEntitiy := entity.KubePod{
				Name:     podFields[0],
				Restarts: podFields[3],
				Age:      podFields[4],
				Status:   podFields[2],
			}
			entities = append(entities, newEntitiy)
		}
	}

	return &entities
}

func (r *kubePodRepository) GetAll() *[]entity.KubePod {
	entities := make([]entity.KubePod, 0)

	var nameSpacesResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420", "kubectl", "get namespaces")
	namespaceInfos := strings.Split(nameSpacesResult, "\n")

	for _, nameSpace := range namespaceInfos[1:] {
		nameSpaceFields := strings.Fields(nameSpace)
		if nameSpaceFields != nil && len(nameSpaceFields) >= 2 {
			var stringResult = r.sshCommandExecuter.RunSshCommand("192.168.55.196:22", "root", "Srvhb0420", "kubectl", "get pods -n "+nameSpaceFields[0])
			podInfos := strings.Split(string(stringResult), "\n")
			for _, podInfo := range podInfos[1:] {
				podFields := strings.Fields(podInfo)
				if podFields != nil && len(podFields) > 2 {
					newEntitiy := entity.KubePod{
						Name:     podFields[0],
						Restarts: podFields[3],
						Age:      podFields[4],
						Status:   podFields[2],
					}
					entities = append(entities, newEntitiy)
				}
			}
		}
	}

	return &entities
}
