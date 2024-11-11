package adapters

import (
	"context"
	"hexagonal_both_publicidad_go/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(uri, dbName, collectionName string) (*MongoRepository, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoRepository{collection: collection}, nil
}

func (repo *MongoRepository) Save(data []domain.Data) error {
	var docs []interface{}
	for _, d := range data {
		docs = append(docs, d)
	}
	_, err := repo.collection.InsertMany(context.TODO(), docs)
	return err
}

func (repo *MongoRepository) IsDuplicate(link string) (bool, error) {
	filter := bson.D{{Key: "link", Value: link}}
	count, err := repo.collection.CountDocuments(context.TODO(), filter)
	return count > 0, err
}
