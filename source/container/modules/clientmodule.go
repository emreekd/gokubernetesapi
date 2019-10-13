package modules

import (
	clientDomain "../../domain/client"
	"../../infrastructure/ssh/client"
)

type ClientModule interface {
	LoadClients() *clientDomain.ISshClient
}

func LoadClients(ip string, username string, password string) clientDomain.ISshClient {
	sshClient := client.InitSshClient(ip, username, password)
	return sshClient
}
