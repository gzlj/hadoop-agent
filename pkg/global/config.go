package global


//import "github.com/gzlj/hadoop-agent/pkg/infra"

// 程序配置
type Config struct {
	Cluster string   `json:"cluster"`
	Master string   `json:"master"`
	ServerPort string   `json:"serverPort"`
	HostName string   `json:"hostName"`

	Phone         string   `json:"phone"`
	AllChannels   []string `json:"allChannels"`
	UsingChannels []string `json:"usingChannels"`
	AllMsgTypes   []string `json:"allMsgTypes"`
	ActiveMsgType []string `json:"activeMsgType"`

}

var (
	// 单例
	G_config *Config
)


func InitConfig(master string, serverPort, hostname, cluster string) (err error) {


	conf := Config{
		Cluster: cluster,
		Master: master,
		ServerPort: serverPort,
		HostName : hostname,
	}
	G_config = &conf
	return
}