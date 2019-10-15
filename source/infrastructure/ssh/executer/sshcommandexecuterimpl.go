package executer

import (
	"bytes"
	"fmt"

	"../../../contract/request"
	builderDomain "../../../domain/builder"
	clientDomain "../../../domain/client"
	executerDomain "../../../domain/executer"
)

type SshCommandExecuter struct {
	sshCommandBuilder builderDomain.ISshCommandBuilder
	sshClient         clientDomain.ISshClient
}

func InitSshExecuter(cb builderDomain.ISshCommandBuilder, sc clientDomain.ISshClient) executerDomain.ISshCommandExecuter {
	executer := &SshCommandExecuter{
		sshCommandBuilder: cb,
		sshClient:         sc,
	}
	return executer
}

func (e *SshCommandExecuter) RunSshCommand(host string, username string, password string, cmdName string, cmd string) string {
	buildRequest := &request.SshCommandBuildRequest{
		CommandName: cmdName,
		Command:     cmd,
	}
	var execCommand = e.sshCommandBuilder.BuildSshCommand(*buildRequest)

	client := e.sshClient.GetConnection(host, username, password)
	sshSession, err := client.NewSession()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer sshSession.Close()
	defer client.Close()

	var stdoutBuf bytes.Buffer
	sshSession.Stdout = &stdoutBuf
	sshSession.Run(execCommand)
	return stdoutBuf.String()
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
