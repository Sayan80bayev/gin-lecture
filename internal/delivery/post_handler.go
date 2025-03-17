package delivery

import (
	"gin-lecture/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.service.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Content     string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.service.Create(req.Title, req.Description, req.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Content     string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Update(req.Title, req.Description, req.Content, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, "Successfully updated post")
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
	}
	c.JSON(http.StatusAccepted, "Successfully deleted post")
}
