package modules

import (
	builderDomain "../../domain/builder"
	"../../infrastructure/ssh/builder"
)

type BuilderModule interface {
	LoadBuilders() *builderDomain.ISshCommandBuilder
}

func LoadBuilders() builderDomain.ISshCommandBuilder {
	sshBuilder := builder.InitSshCommandBuilder()
	return sshBuilder
}
