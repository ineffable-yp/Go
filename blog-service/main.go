package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//路由注册 /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "yinpan"})
	})
	r.Run()
}
