package services

import (
	"hadithgo/domain/entities"
	"hadithgo/domain/models"
)

type HadithService interface {
	ListBooks() (responses []models.GetBookListResponses, err error)
	GetHadith(book string, number int32) (hadith entities.Hadith, err error)
}
