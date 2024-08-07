package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"progress-service/storage/mongodb"
)

type Storage interface {
	NewProgressStorage() Progress
}

func NewMongoStorage(db *mongo.Database) Storage {
	return &StorageImpl{db}
}

type StorageImpl struct {
	db *mongo.Database
}

func (s *StorageImpl) NewProgressStorage() Progress {
	return mongodb.NewProgressRepo(s.db.Collection("progress"))
}
