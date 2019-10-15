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
