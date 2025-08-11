package routes

import (
	"database/sql"
	"github.com/chetan-sharma-exe/sarkari-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	jobHandlers := handlers.NewJobHandler(db)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	jobsPath := router.Group("/api/jobs")
	jobsPath.POST("/createJob", jobHandlers.CreateJob)
	jobsPath.GET("/getAllJobs", jobHandlers.GetAllJobs)
}
