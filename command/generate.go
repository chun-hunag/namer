package command

import (
	"example/name-generator/facade"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"strconv"
)

var (
	GenerateCommand = cli.Command{
		Name:  "generate",
		Usage: "產生假名",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "amount",
				Value: "10",
				Usage: "產生數量",
			},
			cli.StringFlag{
				Name:  "gender",
				Usage: "性別，可用參數： male, female",
			},
			cli.StringFlag{
				Name:  "first-name-num",
				Value: "2",
				Usage: "名的數量， 0 為 1 ~ 2 的亂數，單名給 1 ，雙名給 2",
			},
			cli.BoolFlag{
				Name:  "first-name-only",
				Usage: "只產生名子",
			},
		},
		Action: func(c *cli.Context) error {
			num, err := strconv.Atoi(c.String("amount"))
			if err != nil {
				fmt.Println(errors.Cause(err))
				return nil
			}

			fmt.Println("Generate " + strconv.Itoa(num))
			return facade.Generate(c.GlobalString("provider"), num, func(item string, index int) {
				fmt.Println(item)
			})
		},
	}
)
