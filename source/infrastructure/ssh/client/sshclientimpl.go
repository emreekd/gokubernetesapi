package client

import (
	"fmt"

	sshClientDomain "../../../domain/client"
	"golang.org/x/crypto/ssh"
)

type SshClient struct {
}

func InitSshClient() sshClientDomain.ISshClient {

	sshClient := &SshClient{}
	return sshClient
}

func (sc *SshClient) GetConnection(host string, username string, password string) ssh.Client {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshConn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return *sshConn
}
