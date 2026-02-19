package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"note_api/internal/models"
	"note_api/internal/repository"
)

type BlogHandler struct {
	repo *repository.BlogRepository
}

func NewBlogHandler(repo *repository.BlogRepository) *BlogHandler {
	return &BlogHandler{repo: repo}
}

func (h *BlogHandler) Create(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create"})
		return
	}

	c.JSON(http.StatusCreated, blog)
}

func (h *BlogHandler) GetAll(c *gin.Context) {
	blogs, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (h *BlogHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	blog, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Update(id, &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *BlogHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
