package executer

import (
	executerDomain "../../../domain/executer"
)

type SshCommandExecuter struct {
}

func InitSshExecuter() executerDomain.ISshCommandExecuter {
	executer := &SshCommandExecuter{}
	return executer
}

func (e *SshCommandExecuter) Execute(cmd string) string {
	return ""
}
