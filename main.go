package main

import (
	"os"

	"github.com/JamieYee/save_service/cmd"
	"github.com/urfave/cli/v2"
)

// Usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "v1.0.0"

// @title save_service
// @version v1.0.0
// @description 一个基于 gin-admin 的记账服务后台。
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
func main() {
	app := cli.NewApp()
	app.Name = "save_service"
	app.Version = VERSION
	app.Usage = "一个基于 gin-admin 的记账服务后台。"
	app.Commands = []*cli.Command{
		cmd.StartCmd(),
		cmd.StopCmd(),
		cmd.VersionCmd(VERSION),
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
