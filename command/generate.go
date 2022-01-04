package command

import (
	"example/name-generator/provider"
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

var (
	GenerateCommand = cli.Command{
		Name:  "generate",
		Usage: "產生假名",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "num",
				Value: "10",
				Usage: "產生數量",
			},
		},
		Action: func(c *cli.Context) error {
			return generate(c)
		},
	}
)

func generate(c *cli.Context) error {
	num, err := strconv.Atoi(c.String("num"))

	if err != nil {
		return err
	}

	res, _ := provider.ParseFile("names.yml")
	generator := provider.Create()
	generator.Resource = res

	for i := 0; i < num; i++ {
		fmt.Println(generator.Name())
	}

	return nil
}
