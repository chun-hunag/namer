package command

import (
	"example/name-generator/facade"
	"fmt"
	"github.com/urfave/cli"
)

var (
	QueryCommand = cli.Command{
		Name:  "query",
		Usage: "查詢字典",
		Action: func(c *cli.Context) error {
			str := c.Args().First()

			return facade.Query(str, func(item string) {
				fmt.Println(item)
			})
		},
	}
)
