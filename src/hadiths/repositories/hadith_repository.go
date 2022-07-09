package repositories

import "hadithgo/domain/entities"

type HadithRepository interface {
	FindAll() (hadiths []entities.ListBooks, err error)
}
