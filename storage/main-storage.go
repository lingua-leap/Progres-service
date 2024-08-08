package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"progress-service/storage/mongodb"
)

type Storage interface {
	NewMongoStorage() Mongodb
}

func NewStorage(db *mongo.Database) Storage {
	return &StorageImpl{db}
}

type StorageImpl struct {
	db *mongo.Database
}

func (s *StorageImpl) NewMongoStorage() Mongodb {
	return mongodb.NewProgressRepo(s.db.Collection("progress"))
}
