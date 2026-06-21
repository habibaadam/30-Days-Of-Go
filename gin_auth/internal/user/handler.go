package user

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

type Handler struct {
	service *Service
}


func NewHandler(svc *Service) *Handler {
	return &Handler{
		service:svc,
	}
}

func (h *Handler) Register(c *gin.Context){
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid json body",
		})
		return
	}

	output, err := h.service.Register(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, output)
}