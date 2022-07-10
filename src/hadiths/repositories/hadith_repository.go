package repositories

import "hadithgo/domain/entities"

type HadithRepository interface {
	Get(book string, number int32) (hadith entities.Hadith, err error)
	Insert(hadith entities.InsertHadith) (book entities.Book, err error)
	BulkInsert(book string, path string) (string, error)
}
