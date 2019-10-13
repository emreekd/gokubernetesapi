package client

import (
	sshClientDomain "../../../domain/client"
	"golang.org/x/crypto/ssh"
)

type SshClient struct {
	connection ssh.Client
}

func InitSshClient(ip string, username string, password string) sshClientDomain.ISshClient {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshConn, _ := ssh.Dial("tcp", ip, config)
	defer sshConn.Close()

	sshClient := &SshClient{
		connection: *sshConn,
	}
	return sshClient
}
