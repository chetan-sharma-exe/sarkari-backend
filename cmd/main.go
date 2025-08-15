package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chetan-sharma-exe/sarkari-backend/internal/config"
	"github.com/chetan-sharma-exe/sarkari-backend/internal/db"
	"github.com/chetan-sharma-exe/sarkari-backend/internal/handlers"
	repository "github.com/chetan-sharma-exe/sarkari-backend/internal/repositories"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	 // Allow all origins
	r.Use(cors.New(cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
	}))

	db, err := db.ConnectDatabase(conn.DBUrl)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	h := handlers.NewHandler(repo)

	r.GET("/getAllData", h.GetAll)
	r.POST("/postAllData", h.PostAll)

	if conn.Port == "" {
		conn.Port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", conn.Port))
}
