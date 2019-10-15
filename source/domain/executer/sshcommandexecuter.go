package executer

type ISshCommandExecuter interface {
	Execute(cmdName string, cmd string) string
	RunSshCommand(host string, username string, password string, cmdName string, cmd string) string
}
