package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
	Pinned bool `bson:"pinned" json:"pinned"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type CreateNoteReq struct {
	Title string `json:"title" binding:"required"` // required fields in request
	Content string `json:"content" binding:"required"`
	Pinned bool `json:"pinned"`

}