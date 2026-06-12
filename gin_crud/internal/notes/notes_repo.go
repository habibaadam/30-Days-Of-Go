package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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

func (r *NotesRepo) List(ctx context.Context) ([]Note, error) {
	childContext, cancel := context.WithTimeout(ctx, 5 *time.Second)
	defer cancel()

	filter := bson.M{} // an empty find query filter -> return everything

	cursor, error := r.collection.Find(childContext, filter) // returns a pointer positioned at the first result
	if error != nil {
		return []Note{}, fmt.Errorf("Could not get all notes: %w", error)
	}

	defer cursor.Close(childContext) // close cursor(open connection to mongodb)

	var allNotes []Note

	if err := cursor.All(childContext, &allNotes); err != nil {
		return nil, fmt.Errorf("Getting notes failed: %w", err)
	}
	return allNotes, nil
}