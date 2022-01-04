package command

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

var (
	StatusCommand = cli.Command{
		Name:  "status",
		Usage: "狀態",
		Action: func(c *cli.Context) error {
			readBytes, _ := ioutil.ReadFile("names.yml")

			fmt.Println(`------ File Content Start ------`)
			fmt.Printf("%s", readBytes)
			fmt.Println()
			fmt.Println(`------- File Content End -------`)

			return nil
		},
	}
)
