package configs

import (
	"example/web-service-gin/dtos"
	"example/web-service-gin/helpers"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
	"example/web-service-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(bookRepo *repositories.BookRepository) *gin.Engine {
	route := gin.Default()

	route.POST("/", func(ct *gin.Context) {
		var book models.Book
		err := ct.ShouldBindJSON(&book)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			ct.JSON(http.StatusBadRequest, response)

			return
		}

		status := http.StatusOK
		respon := services.CreateBook(&book, *bookRepo)

		if !respon.Success {
			status = http.StatusBadRequest
		}

		ct.JSON(status, respon)
	})

	route.GET("/", func(c *gin.Context) {
		status := http.StatusOK
		response := services.GetAllBooks(*bookRepo)

		if !response.Success {
			status = http.StatusBadRequest
		}

		c.JSON(status, response)
	})
	route.GET("/:publisher", func(c *gin.Context) {
		publisher := c.Param("publisher")
		status := http.StatusOK
		response := services.GetBookByPublisher(publisher, *bookRepo)

		if !response.Success {
			status = http.StatusBadRequest
		}
		c.JSON(status, response)
	})

	route.PUT("/:publisher", func(c *gin.Context) {
		publisher := c.Param("publisher")
		var book models.Book
		err := c.ShouldBindJSON(&book)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			c.JSON(http.StatusBadRequest, response)
			return
		}

		status := http.StatusOK
		response := services.UpdateBookByPublisher(publisher, &book, *bookRepo)

		if !response.Success {
			status = http.StatusBadRequest
		}
		c.JSON(status, response)
	})

	route.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		status := http.StatusOK
		response := services.DeleteOneBookById(id, *bookRepo)

		if !response.Success {
			status = http.StatusBadRequest
		}
		c.JSON(status, response)
	})

	route.POST("/delete", func(c *gin.Context) {
		var multip dtos.Multi

		err := c.ShouldBindJSON(multip)

		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			c.JSON(http.StatusBadRequest, response)

			return
		}
		if len(multip.Publishers) == 0 {
			response := dtos.Response{Success: false, Message: "Please input your Publisher (min. 1)"}

			c.JSON(http.StatusBadRequest, response)

			return
		}
		status := http.StatusOK
		response := services.DeleteBookByPublisher(&multip, *bookRepo)

		if !response.Success {
			status = http.StatusBadRequest
		}
		c.JSON(status, response)
	})
	return route
}
