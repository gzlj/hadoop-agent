package module


type ComponentStatus struct {
	Name string   `json:"name"`
	Status string   `json:"status"`
}

type ClusteredComponentStatuses struct {
	Cluster           string   `json:"cluster"`
	Hostname          string   `json:"hostname"`
	RunningComponents []string `json:"runningComponents"`
}
