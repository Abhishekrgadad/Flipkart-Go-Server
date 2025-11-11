package users

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewUserStore(db *mongo.Database) (*UserStore, error) {
	col := db.Collection("users")
	store := &UserStore{Col: col}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("idx_email"),
	}
	phoneIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "phone", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("idx_phone"),
	}
	_, err := col.Indexes().CreateMany([]mongo.IndexModel{emailIndex, phoneIndex})
	if err != nil {
		return nil, err
	}
	return store, nil
}
