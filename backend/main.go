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
		var allCompanies = db.GetAllCompanies()
		ctx.JSON(200, gin.H{
			"data": allCompanies,
		})
	})
	r.GET("/person", func(ctx *gin.Context) {
		var allPersons = db.GetAllPersons()
		ctx.JSON(200, gin.H{
			"data": allPersons,
		})
	})

	r.GET("/industry", func(ctx *gin.Context) {
		var allIndustries = db.GetAllIndustries()
		ctx.JSON(200, gin.H{
			"data": allIndustries,
		})
	})

	r.GET("/stock", func(ctx *gin.Context) {
		var allStocks = db.GetAllStocks()
		ctx.JSON(200, gin.H{
			"data": allStocks,
		})
	})

	r.GET("/exchange", func(ctx *gin.Context) {
		var allExchanges = db.GetAllExchange()
		ctx.JSON(200, gin.H{
			"data": allExchanges,
		})
	})
	r.Run()
}
