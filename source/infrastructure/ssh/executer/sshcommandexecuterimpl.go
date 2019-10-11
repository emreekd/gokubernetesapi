package executer

import (
	executerDomain "../../../domain/executer"
)

type SshCommandExecuter struct {
}

func InitSshExecuter() executerDomain.ISshCommandExecuter {
	executer := &executerDomain.ISshCommandExecuter{}
	return executer
}

func (e *SshCommandExecuter) Execute(cmd string) string {
	return ""
}
