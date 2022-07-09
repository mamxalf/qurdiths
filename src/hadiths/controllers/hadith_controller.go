package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hadithgo/domain/models"
	"hadithgo/src/hadiths/services"
	"strconv"
)

type HadithController struct {
	HadithService services.HadithService
}

func NewHadithController(hadithService *services.HadithService) HadithController {
	return HadithController{
		HadithService: *hadithService,
	}
}

func (controller *HadithController) Route(app fiber.Router) {
	app.Get("/books", controller.Books)
	app.Get("/book/:name/:number", controller.Hadith)
}

func (controller *HadithController) Books(c *fiber.Ctx) error {
	responses, _ := controller.HadithService.ListBooks()
	return c.JSON(models.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Success",
		Data:   responses,
	})
}

func (controller *HadithController) Hadith(c *fiber.Ctx) error {
	number, _ := strconv.Atoi(c.Params("number"))
	response, err := controller.HadithService.GetHadith(c.Params("name"), int32(number))

	if err != nil {
		return c.JSON(models.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Error",
			Data:   err.Error(),
		})
	}

	return c.JSON(models.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Success",
		Data:   response,
	})
}
