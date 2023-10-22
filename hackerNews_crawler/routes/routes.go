package routes

import (
	"net/http"

	"github.com/chukwuka-emi/hackerNews_comments_sentiments/hackerNews_crawler/handlers"
	"github.com/gin-gonic/gin"
)

// InitRoutes ...
func InitRoutes(route *gin.Engine){
  
	route.GET("/health",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"Server Okay",
		})
	})

	route.POST("/news-stream",func(ctx *gin.Context) {
		go handlers.CrawlNews()

		ctx.JSON(http.StatusOK,gin.H{
			"message":"Data stream triggered successfuly",
		})
	})
}