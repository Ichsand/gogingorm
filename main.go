package main

import (
	"example/web-service-gin/configs"
	"example/web-service-gin/database"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
	"log"
)

func main() {
	dbUser, dbPassword, dbName := "root", "1234", "gorm_crud_lms"
	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.DB().Ping()

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	defer db.Close()

	bookRepo := repositories.NewBookRepo(db)

	route := configs.SetupRoutes(bookRepo)

	route.Run(":8081")
}
