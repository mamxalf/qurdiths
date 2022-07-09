package services

import "hadithgo/domain/models"

type HadithService interface {
	ListBooks() (responses []models.GetBookListResponses, err error)
}
