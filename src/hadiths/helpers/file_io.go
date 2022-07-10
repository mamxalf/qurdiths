package helpers

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"hadithgo/domain/entities"
	"io/ioutil"
	"os"
)

func BuildBulkInsertData(book string) ([]interface{}, error) {
	// buka file
	rootPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/data/%s.json", rootPath, book)
	var jsonFile, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	var results []entities.HadithJson
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		return nil, err
	}

	var hadiths []interface{}
	for _, result := range results {
		hadiths = append(hadiths, bson.M{
			"book":   book,
			"number": result.Number,
			"arab":   result.Arabian,
			"id":     result.Indonesian,
		})
	}

	return hadiths, nil
}

func ReadFile(book string, number int32) (entities.Hadith, error) {
	// buka file
	rootPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/data/%s.json", rootPath, book)
	var jsonFile, err = os.Open(path)
	if err != nil {
		return entities.Hadith{}, err
	}

	defer jsonFile.Close()

	var results []entities.HadithJson
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		return entities.Hadith{}, err
	}

	var hadith entities.Hadith
	for _, result := range results {
		if result.Number == number {
			hadith = entities.Hadith{
				Book:       book,
				Number:     result.Number,
				Arabian:    result.Arabian,
				Indonesian: result.Indonesian,
			}
			break
		}
	}

	return hadith, nil
}
