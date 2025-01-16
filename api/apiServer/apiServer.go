package apiServer

import (
	"fmt"
	"net/http"

	"github.com/BassicallyDrums/filmbase/api/tmdbConnector"
	"github.com/gin-gonic/gin"
)

func getFilmsJson(c *gin.Context) {
	filmsList, err := tmdbConnector.GetFilms()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, filmsList.Results)
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

func RunAPIServer() error {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/api/films", getFilmsJson)

	return r.Run(":8080")
}
