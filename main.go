package main

import (
	"cmp"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	sloggin "github.com/samber/slog-gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(sloggin.New(logger))
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	port := cmp.Or(os.Getenv("PORT"), "8080")

	logger.Info("Server starting", slog.String("port", port))

	r.Run((":" + port))
}
