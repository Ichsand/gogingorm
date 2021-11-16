package repositories

import (
	"example/web-service-gin/models"

	"github.com/jinzhu/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Save(book *models.Book) Repores {
	err := r.db.Save(book).Error

	if err != nil {
		return Repores{Error: err}
	}

	return Repores{Result: book}
}

func (r *BookRepository) FindAll() Repores {
	var books models.Books

	err := r.db.Find(&books).Error

	if err != nil {
		return Repores{Error: err}
	}

	return Repores{Result: &books}
}

func (r *BookRepository) FindBookById(id string) Repores {
	var book models.Book

	err := r.db.Where(&models.Book{ID: id}).Take(&book).Error

	if err != nil {
		return Repores{Error: err}
	}

	return Repores{Result: &book}
}

func (r *BookRepository) FindBookByPublisher(publisher string) Repores {
	var book models.Book

	err := r.db.Where(&models.Book{Publisher: publisher}).Take(&book).Error

	if err != nil {
		return Repores{Error: err}
	}

	return Repores{Result: &book}
}

func (r *BookRepository) DeleteOneByPublisher(publisher string) Repores {
	err := r.db.Delete(&models.Book{Publisher: publisher}).Error

	if err != nil {
		return Repores{Error: err}
	}

	return Repores{Result: nil}
}

func (r *BookRepository) DeleteOneById(id string) Repores {
	err := r.db.Delete(&models.Book{ID: id}).Error

	if err != nil {
		return Repores{Error: err}
	}

	return Repores{Result: nil}
}

func (r *BookRepository) DeleteByPublishers(publishers *[]string) Repores {
	err := r.db.Where("Publisher IN (?)", *publishers).Delete(&models.Books{}).Error

	if err != nil {
		return Repores{Error: err}
	}

	return Repores{Result: nil}
}
