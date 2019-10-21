package response

type GetNodesResponse struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	InternalIp string `json:"value"`
	Label      string `json:"label"`
}
