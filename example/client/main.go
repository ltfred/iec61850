package main

import (
	"fmt"
	"github.com/ltfred/iec61850"
	"github.com/ltfred/iec61850/example/client/config"
	"github.com/ltfred/iec61850/example/client/worker"
)

func main() {
	address := fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
	subAddress := ""
	if config.SubServerHost != "" && config.SubServerPort != 0 {
		subAddress = fmt.Sprintf("%s:%d", config.SubServerHost, config.ServerPort)
	}
	client := iec61850.NewClient(address, config.Logger, subAddress)
	client.Run(worker.Task)
}
