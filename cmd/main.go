package main

import (
	"fmt"
	"github.com/chetan-sharma-exe/sarkari-backend/internal/config"
	"github.com/chetan-sharma-exe/sarkari-backend/internal/db"
	"github.com/chetan-sharma-exe/sarkari-backend/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db, err := db.ConnectDatabase(conn.DBUrl)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()
	routes.RegisterRoutes(r, db)

	if conn.Port == "" {
		conn.Port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", conn.Port))
}
