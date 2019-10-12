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

func (r *kubePodRepository) GetAll() *[]entity.KubePod {
	var stringResult = r.sshCommandExecuter.Execute("kubectl get pod -n namespace")

	podInfos := strings.Split(string(stringResult), "\n")
	entities := make([]entity.KubePod, 0)

	for _, podInfo := range podInfos {
		podFields := strings.Fields(podInfo)
		if podFields != nil && len(podFields) > 2 {
			newEntitiy := entity.KubePod{
				Name:     podFields[2],
				Restarts: podFields[0],
				Age:      podFields[1],
				Status:   podFields[3],
			}
			entities = append(entities, newEntitiy)
		}
	}

	return &entities
}
