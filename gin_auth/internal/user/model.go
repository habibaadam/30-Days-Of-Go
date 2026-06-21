package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)


type User struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string `bson:"email" json:"email"`
	PasswordHash string `bson:"PasswordHash" json:"-"`
	Role string `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"` // "- means to never include in json output"
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}


type PublicUserRes struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToPublic(u User) PublicUserRes{
	return PublicUserRes{
		ID: u.ID.Hex(),
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,

	}
}