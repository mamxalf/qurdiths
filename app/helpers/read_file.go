package helpers

import (
	"encoding/json"
	"fmt"
	"hadithgo/domain/entities"
	"io/ioutil"
	"os"
)

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
