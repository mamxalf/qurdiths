package repositories

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hadithgo/app/config"
	"hadithgo/app/exceptions"
	"hadithgo/domain/entities"
	"hadithgo/src/hadiths/helpers"
)

type hadithRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewHadithRepository(database *mongo.Database) HadithRepository {
	return &hadithRepositoryImpl{
		Collection: database.Collection("hadiths"),
	}
}
func (repository *hadithRepositoryImpl) Get(book string, number int32) (hadith entities.Hadith, err error) {
	response, err := helpers.ReadFile(book, number)
	fmt.Println(err)
	return response, err
}

func (repository *hadithRepositoryImpl) BulkInsert(book string) (message string, err error) {
	ctx, cancel := config.NewMongoContext(30)
	defer cancel()

	data, _ := helpers.BuildBulkInsertData(book)
	_, err = repository.Collection.InsertMany(ctx, data)
	return book, err
}

// TODO: Insert Books
func (repository *hadithRepositoryImpl) Insert(hadith entities.InsertHadith) (book entities.Book, err error) {
	ctx, cancel := config.NewMongoContext(10)
	defer cancel()

	_, err = repository.Collection.InsertOne(ctx, bson.M{
		"name":   hadith.Name,
		"amount": "amount",
		"file":   "path file",
	})

	exceptions.PanicIfNeeded(err)
	return book, err
}
