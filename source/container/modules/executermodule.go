package modules

import (
	builderDomain "../../domain/builder"
	clientDomain "../../domain/client"
	executerDomain "../../domain/executer"
	"../../infrastructure/ssh/executer"
)

type ExecuterDomain interface {
	LoadExecuters(cb builderDomain.ISshCommandBuilder, sc clientDomain.ISshClient) *executerDomain.ISshCommandExecuter
}

func LoadExecuters(cb builderDomain.ISshCommandBuilder, sc clientDomain.ISshClient) executerDomain.ISshCommandExecuter {
	sshExecuter := executer.InitSshExecuter(cb, sc)
	return sshExecuter
}
