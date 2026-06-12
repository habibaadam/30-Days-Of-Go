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

func (r *NotesRepo) ListID(ctx context.Context, noteId primitive.ObjectID) (Note, error) {
	childContext, cancel := context.WithTimeout(ctx, 5 *time.Second)
	defer cancel()

	filter := bson.M{
		"_id": noteId,
	}

	var note Note

	err := r.collection.FindOne(childContext, filter, options.FindOne()).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("Failed to find the note by the id: %w", err)
	}
	return note, nil
}

func (r *NotesRepo) UpdateById(ctx context.Context, noteId primitive.ObjectID, req UpdateNoteReq) (Note, error) {
	childContext, cancel := context.WithTimeout(ctx, 5 *time.Second)
	defer cancel()

	filter := bson.M{
		"_id": noteId,
	}

	updateValues := bson.M{
		"$set" : bson.M{
			"title": req.Title,
			"content": req.Content,
			"pinned": req.Pinned,
			"UpdatedAt": time.Now().UTC(),
		},
	}

	// additionally configuring findoneandupdate to return the doc after updating the doc.
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedNote Note

	err := r.collection.FindOneAndUpdate(childContext, filter, updateValues, opts).Decode(&updatedNote)
	if err != nil {
		return Note{}, fmt.Errorf("Failed to update the note by the id: %w", err)
	}
	return updatedNote, nil;
}

func (r *NotesRepo) DeletebyId(ctx context.Context, noteId primitive.ObjectID) (bool, error) {
	childContext, cancel := context.WithTimeout(ctx, 5 *time.Second)
	defer cancel()

	filter := bson.M{
		"_id": noteId,
	}

	result, err := r.collection.DeleteOne(childContext, filter)
	if err != nil {
		return false, fmt.Errorf("Failed to delete the note with specified ID: %w", err)
	}

	if (result.DeletedCount == 0) {
		return false, nil
	}
	return true, nil
}