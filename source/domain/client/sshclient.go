package client

import (
	"golang.org/x/crypto/ssh"
)

type ISshClient interface {
	GetConnection(host string, username string, password string) ssh.Client
}
