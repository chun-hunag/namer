package command

import (
	"example/name-generator/facade"
	"example/name-generator/provider"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"strconv"
)

var (
	ServeCommand = cli.Command{
		Name:  "serve",
		Usage: "Start Sever",
		Action: func(c *cli.Context) error {
			return serve(c)
		},
	}
)

func serve(c *cli.Context) error {
	server := gin.Default()
	server.GET(`/generate`, func(g *gin.Context) {
		num, err := strconv.Atoi(g.DefaultQuery("amount", "0"))

		if err != nil {
			g.JSON(400, err)
			return
		}

		firstNameNum, err := strconv.Atoi(g.DefaultQuery("firstNameNum", "0"))

		if err != nil {
			g.JSON(400, err)
			return
		}

		facade.GenerateFirstNameNum = firstNameNum
		facade.GenerateGender = g.DefaultQuery("gender", "")

		str := make([]string, num)
		facade.Generate(c.GlobalString("provider"), num, func(item string, index int) {
			str[index] = item
		})

		provider.Create()
		g.JSON(200, str)
	})

	server.Run()
	return nil
}
