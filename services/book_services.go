package services

import (
	"example/web-service-gin/dtos"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
	"log"

	"github.com/google/uuid"
)

func CreateBook(book *models.Book, repo repositories.BookRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	book.ID = uuidResult.String()

	operationResult := repo.Save(book)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Book)

	return dtos.Response{Success: true, Data: data}
}

func GetAllBooks(bookRepo repositories.BookRepository) dtos.Response {
	operationResult := bookRepo.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	var datas = operationResult.Result.(*models.Books)

	return dtos.Response{Success: true, Data: datas}
}

func GetBookByPublisher(publisher string, bookRepo repositories.BookRepository) dtos.Response {
	operationResult := bookRepo.FindBookByPublisher(publisher)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	var data = operationResult.Result.(*models.Book)

	return dtos.Response{Success: true, Data: data}
}

func GetBookById(id string, bookRepo repositories.BookRepository) dtos.Response {
	operationResult := bookRepo.FindBookById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	var data = operationResult.Result.(*models.Book)

	return dtos.Response{Success: true, Data: data}
}

func UpdateBookByPublisher(publisher string, book *models.Book, bookRepo repositories.BookRepository) dtos.Response {
	existingBookResponse := GetBookByPublisher(publisher, bookRepo)
	if !existingBookResponse.Success {
		return existingBookResponse
	}

	existingBook := existingBookResponse.Data.(*models.Book)

	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Publisher = book.Publisher
	existingBook.Year = book.Year

	operationResult := bookRepo.Save(existingBook)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneBookByPublisher(publisher string, bookRepo repositories.BookRepository) dtos.Response {
	operationResult := bookRepo.DeleteOneByPublisher(publisher)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	return dtos.Response{Success: true}
}

func DeleteOneBookById(id string, bookRepo repositories.BookRepository) dtos.Response {
	operationResult := bookRepo.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	return dtos.Response{Success: true}
}

func DeleteBookByPublisher(multip *dtos.Multi, bookRepo repositories.BookRepository) dtos.Response {
	operationResult := bookRepo.DeleteByPublishers(&multip.Publishers)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	return dtos.Response{Success: true}
}
