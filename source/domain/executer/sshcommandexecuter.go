package executer

type ISshCommandExecuter interface {
	Execute(cmdName string, cmd string) string
}
