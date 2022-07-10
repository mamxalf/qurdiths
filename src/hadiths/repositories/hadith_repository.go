package repositories

import "hadithgo/domain/entities"

type HadithRepository interface {
	FindAll() (hadiths []entities.Book, err error)
	Get(book string, number int32) (hadith entities.Hadith, err error)
	Insert(hadith entities.InsertHadith) (book entities.Book, err error)
}
