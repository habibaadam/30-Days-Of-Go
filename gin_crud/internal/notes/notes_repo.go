package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NotesRepo struct {
	collection *mongo.Collection
}

func NewNotesRepo(db *mongo.Database) *NotesRepo{
	return &NotesRepo{
		collection : db.Collection("notes"),
	}
}

func (r *NotesRepo) Create(ctx context.Context, note Note) (Note, error) {
	childContext, cancel := context.WithTimeout(ctx, 5 *time.Second)
	defer cancel()

	_, error := r.collection.InsertOne(childContext, note)

	if error != nil {
		return Note{}, fmt.Errorf("Insertion of new note failed: %w", error)
	}

	return note, nil
}