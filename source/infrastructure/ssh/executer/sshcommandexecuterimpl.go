package executer

import (
	"strings"

	"../../../contract/request"
	builderDomain "../../../domain/builder"
	executerDomain "../../../domain/executer"
)

type SshCommandExecuter struct {
	sshCommandBuilder builderDomain.ISshCommandBuilder
}

func InitSshExecuter(cb builderDomain.ISshCommandBuilder) executerDomain.ISshCommandExecuter {
	executer := &SshCommandExecuter{
		sshCommandBuilder: cb,
	}
	return executer
}

func (e *SshCommandExecuter) Execute(cmd string) string {
	buildRequest := &request.SshCommandBuildRequest{
		CommandName: "ls",
		Command:     "-lh",
	}
	var execCommand = e.sshCommandBuilder.Build(*buildRequest)

	stdout, err := execCommand.Output()

	if err != nil {
		println(err.Error())
		return err.Error()
	}
	podInfos := strings.Split(string(stdout), "\n")

	for _, podInfo := range podInfos {
		podFields := strings.Fields(podInfo)
		if podFields != nil && len(podFields) > 2 {
			println(podFields[2])
		}
	}

	return string(stdout)
}
