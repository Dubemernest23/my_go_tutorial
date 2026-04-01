package notes

import (
	"errors"
	"fmt"
	"net/http"

	// "notes-api/internal/notes"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// holds the dependecies and handlers for the notes routes

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler { return &Handler{repo: repo} }

func (h *Handler) createNote(c *gin.Context) {
	// parse request body into CreateNoteRequest struct
	var req CreateNoteRequest

	// binds json body to the struct and validates it based on the binding tags
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Json format",
		})
		return
	}

	now := time.Now().UTC()

	note := Note{
		ID:        primitive.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	created, err := h.repo.create(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create note",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *Handler) ListNotes(c *gin.Context) {
	notes, err := h.repo.List(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to list notes",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}

func (h *Handler) GetNotesById(c *gin.Context) {

	idString := c.Param("id")

	// convert 24 hex str to mongo objID
	objId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error ":  " Invalid id or id mismatch",
		})
		return
	}
	note, err := h.repo.GetByID(c.Request.Context(), objId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error ":  " Invalid id or id mismatch",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error ":  " Server Error occurred",
		})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *Handler) UpdateNoteById(c *gin.Context) {
	idString := c.Param("id")

	// convert 24 hex str to mongo objID
	objId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error ":  " Invalid id or id mismatch",
		})
		return
	}

	// create request
	var req UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error ":  " Invalid Json format",
		})
		return
	}

	updatedNote, err := h.repo.UpdateByID(c.Request.Context(), objId, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error ":  " Invalid id or id mismatch",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error ":  " Server Error occurred",
		})
		fmt.Printf("Server error %v", err.Error())
		return
	}
	c.JSON(http.StatusOK, updatedNote)

}

// delete note by id
func (h *Handler) DeleteNoteById(c *gin.Context) {
	idString := c.Param("id")

	// convert 24 hex str to mongo objID
	objId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid id or id mismatch",
		})
		return
	}

	err = h.repo.DeleteByID(c.Request.Context(), objId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Note not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Server Error occurred",
		})
		fmt.Printf("Server error %v", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Note deleted successfully",
	})
}
