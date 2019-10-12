package builder

import (
	exec "os/exec"

	"../../../contract/request"
	sshBuilder "../../../domain/builder"
)

type SshCommandBuilder struct {
}

func InitSshCommandBuilder() sshBuilder.ISshCommandBuilder {
	builder := &SshCommandBuilder{}
	return builder
}

func (b *SshCommandBuilder) Build(request request.SshCommandBuildRequest) *exec.Cmd {
	cmd := exec.Command(request.CommandName, request.Command)
	return cmd
}
