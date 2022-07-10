package repositories

import "hadithgo/domain/entities"

type HadithBookRepository interface {
	FindAll() (hadits []entities.Book, err error)
}
