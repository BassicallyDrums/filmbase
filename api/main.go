package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Film struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Rating int    `json:"rating"`
}

var films = []Film{
	{ID: 0, Title: "Gladiator 2", Rating: 3},
	{ID: 1, Title: "The Kind", Rating: 5},
}

func getFilmsJson(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, films)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/api/films", getFilmsJson)

	r.Run(":8080")

}
