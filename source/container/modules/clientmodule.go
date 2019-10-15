package modules

import (
	clientDomain "../../domain/client"
	"../../infrastructure/ssh/client"
)

type ClientModule interface {
	LoadClients() *clientDomain.ISshClient
}

func LoadClients() clientDomain.ISshClient {
	sshClient := client.InitSshClient()
	return sshClient
}
