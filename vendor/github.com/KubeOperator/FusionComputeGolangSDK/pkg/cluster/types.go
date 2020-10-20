package cluster

type Cluster struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
	Urn  string `json:"urn"`
}

type ListClusterResponse struct {
	Clusters []Cluster `json:"clusters"`
}
