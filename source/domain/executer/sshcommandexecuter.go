package executer

type ISshCommandExecuter interface {
	Execute(cmd string) string
}
