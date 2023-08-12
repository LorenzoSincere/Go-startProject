package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"startProject/go-project-example/Repository"
	"startProject/go-project-example/Util"
	"startProject/go-project-example/handler"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(1)
	}
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := handler.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}

}

func Init() error {
	if err := Repository.Init(); err != nil {
		return err
	}
	if err := Util.InitLogger(); err != nil {
		return err
	}
	return nil
}
