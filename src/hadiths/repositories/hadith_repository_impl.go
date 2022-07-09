package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hadithgo/app/config"
	"hadithgo/app/exceptions"
	"hadithgo/domain/entities"
)

type hadithRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewHadithRepository(database *mongo.Database) HadithRepository {
	return &hadithRepositoryImpl{
		Collection: database.Collection("books"),
	}
}

func (repository *hadithRepositoryImpl) FindAll() (hadiths []entities.ListBooks, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exceptions.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exceptions.PanicIfNeeded(err)

	for _, document := range documents {
		hadiths = append(hadiths, entities.ListBooks{
			ID:     document["_id"].(primitive.ObjectID).Hex(),
			Name:   document["name"].(string),
			Amount: document["amount"].(int32),
		})
	}

	return hadiths, err
}
