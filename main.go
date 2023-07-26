package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/q10357/cat-api-graphql/graph"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/cat", graph.CatGraphRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("127.0.0.1:" + port)
}
