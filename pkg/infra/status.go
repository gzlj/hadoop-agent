package infra

import (
	"github.com/gzlj/hadoop-agent/pkg/global"
	"github.com/gzlj/hadoop-agent/pkg/module"
	"strings"
)


//func GetComponentStatus() (statues []module.ComponentStatus, err error){
//ClusteredComponentStatuses
func GetComponentStatus() (status module.ClusteredComponentStatuses, err error){
	var (
		bytes []byte
		names []string
	)

	bytes, err = RunAndOutput(global.HadoopComponentNamesCmd)
	if err != nil {
		return
	}

	names = strings.Split(strings.TrimSpace(string(bytes)), " ")
	if len(names) == 0 {
		return
	}
	for _, n :=  range names {
		if len(n) == 0 {
			continue
		}
		names = append(names, n)
	}
	status = constructComponentStatuses(names)
	return
}

func constructComponentStatuses(components []string) (status module.ClusteredComponentStatuses) {
	status.RunningComponents = components
	status.ClusterName = global.G_config.Cluster
	status.Host =global.G_config.HostName
	return
}




