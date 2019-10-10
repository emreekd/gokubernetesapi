package entity

type KubePod struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Restarts string `json:"restarts"`
	Age      string `json:"age"`
}
