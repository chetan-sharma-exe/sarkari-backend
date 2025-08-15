package handlers

import (
	"net/http"

	"github.com/chetan-sharma-exe/sarkari-backend/internal/models"
	repository "github.com/chetan-sharma-exe/sarkari-backend/internal/repositories"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) GetAll(c *gin.Context) {
	data := h.Repo.GetAllData()
	c.JSON(http.StatusOK, data)
}

func (h *Handler) PostAll(c *gin.Context) {
	var data models.AllData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Repo.InsertAllData(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}

