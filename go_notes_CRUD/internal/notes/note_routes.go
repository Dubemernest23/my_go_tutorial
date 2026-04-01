package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
	// create the handler and repo once at startup and pass the repo to the handler
	// then register the handler methods as route handlers for the appropriate endpoints

	repo := NewRepo(db)
	handler := NewHandler(repo)

	notesGroup := r.Group("/notes")
	{
		notesGroup.POST("", handler.createNote)
		notesGroup.GET("", handler.ListNotes)
		notesGroup.GET("/:id", handler.GetNotesById)
		notesGroup.PUT("/:id", handler.UpdateNoteById)
		notesGroup.DELETE("/:id", handler.DeleteNoteById)
	}

}
