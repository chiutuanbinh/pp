package main

import (
	"mysite/db"
	"mysite/orm"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.GET("/companies", func(ctx *gin.Context) {
		var companies []orm.Company
		orm.DB.Find(&companies)
		ctx.JSON(200, gin.H{
			"data": companies,
		})
	})
	r.GET("/companies/:id/", func(ctx *gin.Context) {
		cid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		var company orm.Company
		orm.DB.First(&company, cid)
		if company.ID != cid {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(200, gin.H{
			"data": company,
		})
	})

	r.GET("/stock/:code", func(ctx *gin.Context) {
		code := ctx.Param("code")
		var stock orm.Stock
		orm.DB.Where("code=?", code).Find(&stock)
		var stockPrices []orm.StockPrice = make([]orm.StockPrice, 100)

		orm.DB.Limit(100).Order("date desc").Where("stock_id=?", stock.ID).Find(&stockPrices)
		ctx.JSON(200, gin.H{
			"data": stockPrices,
		})
	})
	r.POST("/companies", func(ctx *gin.Context) {
		company := db.Company{}
		ctx.ShouldBindJSON(&company)
		company.AddToDb()
		ctx.JSON(200, gin.H{
			"data": company,
		})
	})
	r.GET("/persons", func(ctx *gin.Context) {
		var allPersons []orm.Person
		orm.DB.Find(&allPersons)
		ctx.JSON(200, gin.H{
			"data": allPersons,
		})
	})
	r.POST("/persons", func(ctx *gin.Context) {
		person := db.Person{}
		ctx.ShouldBindJSON(&person)
		person.AddToDb()
		ctx.JSON(200, gin.H{
			"data": person,
		})
	})
	r.PUT("/persons/:id", func(ctx *gin.Context) {
		person := db.Person{}
		ctx.ShouldBindJSON(&person)
		person.AddToDb()
		ctx.JSON(200, gin.H{
			"data": person,
		})
	})
	r.GET("/persons/:id/", func(ctx *gin.Context) {
		pid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		var person = db.GetPerson(pid)
		if person.ID != pid {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(200, gin.H{
			"data": person,
		})
	})

	r.GET("/industries", func(ctx *gin.Context) {
		var allIndustries = db.GetAllIndustries()
		ctx.JSON(200, gin.H{
			"data": allIndustries,
		})
	})

	r.GET("/industries/:id/", func(ctx *gin.Context) {
		iid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		var industry = db.GetIndustry(iid)
		if industry.ID != iid {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(200, gin.H{
			"data": industry,
		})
	})

	r.GET("/stocks", func(ctx *gin.Context) {
		var allStocks = db.GetAllStocks()
		ctx.JSON(200, gin.H{
			"data": allStocks,
		})
	})

	r.GET("/stocks/:id", func(ctx *gin.Context) {
		sid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		var stock = db.GetStock(sid)
		if stock.ID != sid {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(200, gin.H{
			"data": stock,
		})
	})

	r.POST("/stocks", func(ctx *gin.Context) {

	})

	r.GET("/exchanges", func(ctx *gin.Context) {
		var allExchanges = db.GetAllExchange()
		ctx.JSON(200, gin.H{
			"data": allExchanges,
		})
	})
	r.GET("/exchanges/:id", func(ctx *gin.Context) {
		eid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnprocessableEntity)
			return
		}
		var exchange = db.GetExchange(eid)
		if exchange.ID != eid {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(200, gin.H{
			"data": exchange,
		})
	})
	r.Run()
}
