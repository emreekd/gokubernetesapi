package model

type SshCommandResult struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Restarts string `json:"restarts"`
	Age      string `json:"age"`
}
