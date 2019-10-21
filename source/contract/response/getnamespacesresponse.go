package response

type GetNamespacesResponse struct {
	Namespaces []Namespace `json:"namespaces"`
}
type Namespace struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Age    string `json:"age"`
}
