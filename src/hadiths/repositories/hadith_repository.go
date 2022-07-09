package repositories

import "hadithgo/domain/entities"

type HadithRepository interface {
	FindAll() (hadiths []entities.ListBooks, err error)
	Get(book string, number int32) (hadith entities.Hadith, err error)
}
