package response

type GetAllPodsResponse struct {
	Pods []Pod `json:"pods"`
}

type Pod struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Restarts string `json:"restarts"`
	Age      string `json:"age"`
}
