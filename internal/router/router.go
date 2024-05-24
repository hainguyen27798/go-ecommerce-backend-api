package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
