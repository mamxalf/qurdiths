package repositories

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hadithgo/app/config"
	"hadithgo/app/exceptions"
	"hadithgo/app/helpers"
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

func (repository *hadithRepositoryImpl) FindAll() (hadiths []entities.Book, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exceptions.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exceptions.PanicIfNeeded(err)

	for _, document := range documents {
		hadiths = append(hadiths, entities.Book{
			ID:     document["_id"].(primitive.ObjectID).Hex(),
			Name:   document["name"].(string),
			Amount: document["amount"].(int32),
		})
	}

	return hadiths, err
}

func (repository *hadithRepositoryImpl) Get(book string, number int32) (hadith entities.Hadith, err error) {
	response, err := helpers.ReadFile(book, number)
	fmt.Println(err)
	return response, err
}

// TODO: Insert Books
func (repository *hadithRepositoryImpl) Insert(hadith entities.InsertHadith) (book entities.Book, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err = repository.Collection.InsertOne(ctx, bson.M{
		"name":   hadith.Name,
		"amount": "amount",
		"file":   "path file",
	})

	exceptions.PanicIfNeeded(err)
	return book, err
}
