package main

import (
	"mysite/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.GET("/company", func(ctx *gin.Context) {
		var allCompany = db.GetAllCompany()
		ctx.JSON(200, gin.H{
			"data": allCompany,
		})
	})
	r.GET("")
	r.Run()
}
