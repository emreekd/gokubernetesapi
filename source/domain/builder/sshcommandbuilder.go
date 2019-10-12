package builder

import (
	exec "os/exec"

	"../../contract/request"
)

type ISshCommandBuilder interface {
	Build(Request request.SshCommandBuildRequest) *exec.Cmd
}
