package main

import (
	"example/name-generator/command"
	"github.com/urfave/cli"
	"os"
)

func main() {
	os.Chdir(".") // change current working directory
	app := cli.NewApp()
	app.Name = "Namer"
	app.Commands = []cli.Command{
		command.GenerateCommand,
		command.StatusCommand,
		command.QueryCommand,
		command.ServeCommand,
	}

	app.Run(os.Args)
}
