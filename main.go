package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MaxBodyBytes = 10
)

func main() {
	router := gin.Default()
	router.Use(bodySizeMiddleware)

	router.POST("/", func(ctx *gin.Context) {
		m := map[string]string{}
		if err := ctx.BindJSON(&m); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		ctx.JSON(200, m)
	})

	router.Run()
}

func bodySizeMiddleware(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, MaxBodyBytes)

	c.Next()
}
