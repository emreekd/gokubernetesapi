package entity

type Deployment struct {
	Name          string `json:"name"`
	Ready         string `json:"ready"`
	UpToDate      string `json:"upToDate"`
	Available     string `json:"available"`
	Age           string `json:"age"`
	ContainerName string `json:"containerName"`
}
