package modules

import (
	builderDomain "../../domain/builder"
	executerDomain "../../domain/executer"
	"../../infrastructure/ssh/executer"
)

type ExecuterDomain interface {
	LoadExecuters(cb builderDomain.ISshCommandBuilder) *executerDomain.ISshCommandExecuter
}

func LoadExecuters(cb builderDomain.ISshCommandBuilder) executerDomain.ISshCommandExecuter {
	sshExecuter := executer.InitSshExecuter(cb)
	return sshExecuter
}
