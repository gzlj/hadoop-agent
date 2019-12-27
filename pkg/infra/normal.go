package infra

import (
	"context"
	"os/exec"
	"strings"
)

func RunAndOutput(cmdStr string) (bytes []byte,err error) {
	cmd := exec.CommandContext(context.TODO(), "bash", "-c", cmdStr)
	bytes, err = cmd.Output()
	return
}

func GetHostName() (hostname string, err error) {
	var bytes []byte
	bytes, err = RunAndOutput("hostname")
	if err != nil {
		return
	}
	hostname = strings.TrimSpace(string(bytes))
	return
}
