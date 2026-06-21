package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	collection *mongo.Collection
}

func UserRepo(db *mongo.Database) *Repo {
	return &Repo{
		collection: db.Collection("users"),
	}
}

func (u *Repo) Create(ctx context.Context, user User) (User, error) {

	user.ID = bson.NewObjectID()
	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("Failed to create user: %w", err)
	}

	/*
	id, ok := new_user.InsertedID.(bson.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("Failed to create user and inserted id is not objectid")
	}
	user.ID = id
	*/

	return user, nil
}

func (u *Repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	filter := bson.M{"email" : email}

	var user User

	err := u.collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, mongo.ErrNoDocuments
		}
		return User{}, fmt.Errorf("find by email failed: %w", err)
	}
	return user, nil

}
