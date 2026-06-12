package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	repo *NotesRepo
}

func NewHandler(repo *NotesRepo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote(c *gin.Context) {
	var req CreateNoteReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	now := time.Now().UTC()

	newNote := Note{
		ID: primitive.NewObjectID(),
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	created, err := h.repo.Create(c.Request.Context(), newNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create a note",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": created,
	})
}

func (h *Handler) ListNotes(c *gin.Context) {
	notes, err := h.repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch all notes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"All notes returned": notes,
	})
}