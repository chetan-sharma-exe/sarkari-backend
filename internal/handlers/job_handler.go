package handlers

import (
	"database/sql"
	"net/http"
	m "prem/internal/models"
	"prem/internal/repositories"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	DB *sql.DB
}

func NewJobHandler(db *sql.DB) *JobHandler {
	return &JobHandler{
		DB: db,
	}
}

func (j *JobHandler) CreateJob(c *gin.Context) {
	var job m.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.CreateJob(j.DB, job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "job created"})
}

func (j *JobHandler) GetAllJobs(c *gin.Context) {
	jobs, err := repositories.GetJobs(j.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}
