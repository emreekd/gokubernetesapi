package executer

import (
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

func (e *SshCommandExecuter) Execute(cmdName string, cmd string) string {
	buildRequest := &request.SshCommandBuildRequest{
		CommandName: cmdName,
		Command:     cmd,
	}
	var execCommand = e.sshCommandBuilder.Build(*buildRequest)

	stdout, err := execCommand.Output()

	if err != nil {
		println(err.Error())
		return err.Error()
	}

	return string(stdout)
}
