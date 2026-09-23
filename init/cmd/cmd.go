package cmd

import (
	"fmt"
	"github.com/pangtaek/MyProject/config"
	"github.com/pangtaek/MyProject/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	fmt.Println(c.config.Server.Port)

	c.network.StartServer(c.config.Server.Port)

	return c
}
