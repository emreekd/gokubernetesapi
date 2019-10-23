package persistance

import (
	"fmt"
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
	var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "get pod "+podName+"  -n "+namespace+" -o json")
	return string(stringResult)
}

func (r *kubePodRepository) RestartPod(podName string, namespace string) bool {
	var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "delete pod "+podName+"  -n "+namespace)
	var result = string(stringResult)
	if strings.Contains(result, podName) && strings.Contains(result, "deleted") {
		return true
	}
	return false
}

func (r *kubePodRepository) GetNodes() *[]entity.Node {
	entities := make([]entity.Node, 0)
	var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "get nodes -o jsonpath='{.items[*].status.addresses[?(@.type==\"InternalIP\")].address}'")
	var nodeInfos = strings.Split(string(stringResult), " ")
	for _, node := range nodeInfos {
		newNode := entity.Node{
			InternalIp: node,
		}
		entities = append(entities, newNode)
	}
	return &entities
}

func (r *kubePodRepository) GetNamespaces() *[]entity.Namespace {
	entities := make([]entity.Namespace, 0)
	var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "get namespaces")
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
	var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "get deployments -n "+namespace)
	deploymentInfos := strings.Split(string(stringResult), "\n")
	containerInfoResults := make(chan *deploymentInfoResult, 0)
	channelLength := 0
	for _, deploymentInfo := range deploymentInfos[1:] {
		deploymentFields := strings.Fields(deploymentInfo)
		if deploymentFields != nil && len(deploymentFields) >= 4 {
			channelLength++
			go func() {
				containerInfoResults <- GetDeploymentInfo(r.sshCommandExecuter, deploymentFields, namespace)
			}()
		}
	}
	for i := 0; i < channelLength; i++ {
		containerInfoResult := <-containerInfoResults
		containerInfos := strings.Split(string(containerInfoResult.CommandResult), " ")
		containerName := containerInfos[1]
		imageName := containerInfos[2]
		containerInfoResult.Deployment.ContainerName = containerName
		containerInfoResult.Deployment.Image = imageName

		entities = append(entities, containerInfoResult.Deployment)
	}
	return &entities
}

type deploymentInfoResult struct {
	Deployment    entity.Deployment
	CommandResult string
}

func GetDeploymentInfo(sshCommandExecuter executer.ISshCommandExecuter, deploymentFields []string, namespace string) *deploymentInfoResult {
	cmdResult := sshCommandExecuter.RunSshCommand("host","username","password", "kubectl",
		"get deployment "+deploymentFields[0]+" -n "+namespace+" -o=jsonpath='{.metadata.name}{\" \"}{..containers[0].name}{\" \"}{..containers[0].image}'")

	newEntity := entity.Deployment{
		Name:      deploymentFields[0],
		Ready:     deploymentFields[1],
		UpToDate:  deploymentFields[2],
		Available: deploymentFields[3],
		Age:       deploymentFields[4],
	}

	deploymentInfo := &deploymentInfoResult{
		Deployment:    newEntity,
		CommandResult: cmdResult,
	}

	return deploymentInfo

}

func (r *kubePodRepository) UpdateImageForDeployment(deploymentName string, containerName string, newImage string, namespace string) bool {
	var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",,
		"kubectl", "set image deployment/"+deploymentName+" "+containerName+"="+newImage+" -n "+namespace)
	result := string(stringResult)
	fmt.Println(result)
	if result == "" || strings.Contains(result, "updated") {
		return true
	}
	return false
}

func (r *kubePodRepository) GetByNamespace(namespace string) *[]entity.KubePod {
	entities := make([]entity.KubePod, 0)
	var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "get pods -n "+namespace)
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

	var nameSpacesResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "get namespaces")
	namespaceInfos := strings.Split(nameSpacesResult, "\n")

	for _, nameSpace := range namespaceInfos[1:] {
		nameSpaceFields := strings.Fields(nameSpace)
		if nameSpaceFields != nil && len(nameSpaceFields) >= 2 {
			var stringResult = r.sshCommandExecuter.RunSshCommand("host","username","password",, "kubectl", "get pods -n "+nameSpaceFields[0])
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
