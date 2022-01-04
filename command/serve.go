package command

import (
	"example/name-generator/provider"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
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
	resource, _ := provider.ParseFile("names.yml")
	num := 10

	server := gin.Default()
	server.GET(`/generate`, func(c *gin.Context) {
		generator := provider.Create()
		generator.Resource = resource

		str := []string{}
		for i := 0; i < num; i++ {
			str = append(str, generator.Name())
		}

		provider.Create()
		c.JSON(200, str)
	})

	server.Run()
	return nil
}
