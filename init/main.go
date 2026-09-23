package main

import (
	"flag"
	"github.com/pangtaek/MyProject/init/cmd"
)

var configPathFlag = flag.String("config", "./config.toml", "path to config file")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
 