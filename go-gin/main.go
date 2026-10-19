package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.String(200, "hello from gin") })
	r.Run() // reads $PORT
}
