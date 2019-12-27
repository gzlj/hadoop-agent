package module


type ComponentStatus struct {
	Name string   `json:"name"`
	Status string   `json:"status"`
}

type ClusteredComponentStatuses struct {
	ClusterName string   `json:"clusterName"`
	Host string   `json:"host"`
	RunningComponents []string `json:"runningComponents"`
}
