package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hadithgo/domain/models"
	"hadithgo/src/hadiths/services"
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
	app.Get("/lists", controller.Lists)
}

func (controller *HadithController) Lists(c *fiber.Ctx) error {
	responses, _ := controller.HadithService.ListBooks()
	return c.JSON(models.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Success",
		Data:   responses,
	})
}
