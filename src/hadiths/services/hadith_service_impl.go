package services

import (
	"hadithgo/domain/entities"
	"hadithgo/domain/models"
	"hadithgo/src/hadiths/repositories"
)

type hadithServiceImpl struct {
	HadithRepository repositories.HadithRepository
}

func NewHadithService(hadithRepository *repositories.HadithRepository) HadithService {
	return &hadithServiceImpl{
		HadithRepository: *hadithRepository,
	}
}

func (service *hadithServiceImpl) ListBooks() (responses []models.GetBookListResponses, err error) {
	bookLists, err := service.HadithRepository.FindAll()
	for _, book := range bookLists {
		responses = append(responses, models.GetBookListResponses{
			ID:     book.ID,
			Name:   book.Name,
			Amount: book.Amount,
		})
	}
	return responses, err
}

func (service *hadithServiceImpl) GetHadith(book string, number int32) (hadith entities.Hadith, err error) {
	return service.HadithRepository.Get(book, number)
}
