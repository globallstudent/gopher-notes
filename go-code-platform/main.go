package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-code-platform/judge"
	"go-code-platform/problems"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Serve HTML
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API endpoints
	r.GET("/api/problems", problems.ListProblems)
	r.POST("/api/submit", judge.HandleSubmission)

	r.Run(":8080")
}
