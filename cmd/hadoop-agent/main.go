package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/hadoop-agent/pkg/global"
	"github.com/gzlj/hadoop-agent/pkg/handler"
	"github.com/gzlj/hadoop-agent/pkg/infra"
	"os"
	"runtime"
)

type APIServer struct {
	engine *gin.Engine
	port string
}

func (s *APIServer) Run() {
	s.engine.Run(":" + s.port)
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func init() {
	initEnv()
	master := os.Getenv("MASTER")
	if master == ""{
		master="localhost"
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == ""{
		serverPort="18080"
	}
	cluster := os.Getenv("CLUSTER")
	if cluster == ""{
		cluster="mycluster"
	}
	hostname , _ := infra.GetHostName()
	global.InitConfig(master, serverPort, hostname, cluster)
	go infra.HeartBeat(master)
}

func (s *APIServer) registryApi() {
	registryBasicApis(s.engine)
}

func registryBasicApis(r *gin.Engine) {
	r.GET("/status/components", handler.HandleGetComponentStatuses)
}

func main() {
	server := &APIServer{
		engine: gin.Default(),
		port: global.G_config.ServerPort,
	}
	go infra.HeartBeatLoop()
	server.registryApi()
	server.Run()
}
