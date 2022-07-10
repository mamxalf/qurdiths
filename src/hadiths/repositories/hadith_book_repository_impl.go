package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hadithgo/app/config"
	"hadithgo/app/exceptions"
	"hadithgo/domain/entities"
)

type hadithBookRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewHadithBookRepository(database *mongo.Database) HadithBookRepository {
	return &hadithBookRepositoryImpl{
		Collection: database.Collection("books"),
	}
}

func (repository *hadithBookRepositoryImpl) FindAll() (books []entities.Book, err error) {
	ctx, cancel := config.NewMongoContext(10)
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exceptions.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exceptions.PanicIfNeeded(err)

	for _, document := range documents {
		books = append(books, entities.Book{
			ID:     document["_id"].(primitive.ObjectID).Hex(),
			Name:   document["name"].(string),
			Amount: document["amount"].(int32),
		})
	}

	return books, err
}
