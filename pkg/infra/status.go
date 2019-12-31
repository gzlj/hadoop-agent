package infra

import (
	"encoding/json"
	"errors"
	"github.com/gzlj/hadoop-agent/pkg/global"
	"github.com/gzlj/hadoop-agent/pkg/module"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)


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
	status.Cluster = global.G_config.Cluster
	status.Hostname =global.G_config.HostName
	return
}

//send status to master period
func HeartBeat(server string) (err error) {
	var (
		resp *http.Response
		req  *http.Request
		br *module.BusinessResponse
		body []byte
		status module.ClusteredComponentStatuses
	)
	status, err = GetComponentStatus()
	if err != nil {
		return
	}
	body, err = json.Marshal(status)
	if err != nil {
		return
	}
	req = ConstructHttpPostReq(global.G_config.Master, global.HEART_BEART_URI, string(body))
	resp, err = http.DefaultClient.Do(req)
	if(err != nil){
		err = errors.New("Failed to send heartbeat to master: "+ err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = errors.New("Failed to send heartbeat to master: " + resp.Status)
		return
	}
	br = &module.BusinessResponse{}
	body, _ = ioutil.ReadAll(resp.Body)
	err =json.Unmarshal(body, br)
	if br.Code != 200 {
		err = errors.New("Failed to send heartbeat to master: " + br.Message)
		return
	}
	return
}

func ConstructHttpPostReq(server, uri,body string) (req *http.Request){
	var url string
	url = global.HTTP_PREFIX + server + uri
	req , _ = http.NewRequest("POST", url, strings.NewReader(body))
	return
}

func HeartBeatLoop() {
	var (
		err error
	)
	timer1 := time.NewTimer(30 * time.Second)
	for {
		select {
		case <-timer1.C:
			err = HeartBeat(global.G_config.Master)
			if err != nil {
				log.Println("Error happened when send heartbeat to master: " ,err)
			}
			timer1.Reset(30 * time.Second)
		}
	}

}


