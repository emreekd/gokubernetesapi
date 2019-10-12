package request

type SshCommandBuildRequest struct {
	CommandName string `json:"commandName"`
	Command     string `json:"command"`
}
