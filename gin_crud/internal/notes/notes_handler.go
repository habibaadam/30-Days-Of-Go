package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

	// req body validation
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

func (h *Handler) GetNoteById(c *gin.Context) {

	idStr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid ID Format",
		})
	}
	note, err := h.repo.ListID(c.Request.Context(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments){
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Note based on ID was not found",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Note found details": note,
	})
}

func (h *Handler) UpdateNoteById(c *gin.Context) {

	idStr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid ID Format",
		})
	}

	var req UpdateNoteReq

	// req body validation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	updateNote, err := h.repo.UpdateById(c.Request.Context(), objID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update a note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated note details": updateNote,
	})
}

func (h *Handler) DeleteNoteByID(c *gin.Context) {
	idStr := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid ID Format",
		})
		return
	}

	deletedNote, err := h.repo.DeletebyId(c.Request.Context(), objID)

	if !deletedNote {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note with ID not found",
		})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Failed to delete a note by the given ID",
	})
	return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleted": true,
	})

}