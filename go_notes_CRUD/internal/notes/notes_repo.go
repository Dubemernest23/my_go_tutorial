package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// HTTP Request → Server/Handler → Service (business logic) → Repository → Database
// The repository sits between your business logic and the database, handling all CRUD operations.
// database access logic is abstracted away in the repository, allowing your business logic to focus on
// application-specific rules and workflows.

type Repo struct {
	// db connection or collection reference

	coll *mongo.Collection
}

// create collection reference and return repo instance
func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

// you will have to create methods for CRUD operations here, e.g. CreateNote, GetNoteByID, UpdateNote
func (r *Repo) create(ctx context.Context, note Note) (Note, error) {

	// opCtx is a child context with a timeout for the database operation
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)
	if err != nil {
		return Note{}, fmt.Errorf("Failed to create note: %w", err)
	}

	return note, nil
}

// get all notes
func (r *Repo) List(ctx context.Context) ([]Note, error) {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{} // empty filter to get all documents

	cursor, err := r.coll.Find(opctx, filter)

	if err != nil {
		return nil, fmt.Errorf("Failed to list notes: %w", err)
	}
	defer cursor.Close(opctx) // cursor must close after use to avoid resource leaks

	var notes []Note
	if err := cursor.All(opctx, &notes); err != nil {
		return nil, fmt.Errorf("Failed to decode notes: %w", err)
	}
	return notes, nil
}

// get notes by id
func (r *Repo) GetByID(ctx context.Context, id primitive.ObjectID) (Note, error) {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var note Note
	err := r.coll.FindOne(opctx, filter, options.FindOne()).Decode(&note)

	if err != nil {
		return Note{}, fmt.Errorf("Failed to get note by ID: %w", err)
	}

	return note, nil

}

// update note by id
func (r *Repo) UpdateByID(ctx context.Context, id primitive.ObjectID, req UpdateNoteRequest) (Note, error) {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var updatedNote Note
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"title":      req.Title,
			"content":    req.Content,
			"pinned":     req.Pinned,
			"updated_at": time.Now().UTC(),
		},
	}

	//Use the builder, not a struct literal
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := r.coll.FindOneAndUpdate(opctx, filter, update, opts).Decode(&updatedNote)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Note{}, mongo.ErrNoDocuments // let handler deal with it
		}
		return Note{}, err
	}

	return updatedNote, nil
}

// delete note by id
func (r *Repo) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	opctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	_, err := r.coll.DeleteOne(opctx, filter)
	if err != nil {
		return err
	}

	return nil
}
