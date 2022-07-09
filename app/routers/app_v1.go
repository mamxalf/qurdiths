package routers

import (
	"github.com/gofiber/fiber/v2"
	"hadithgo/app/config"
	"hadithgo/src/hadiths/controllers"
	"hadithgo/src/hadiths/repositories"
	"hadithgo/src/hadiths/services"
)

func SetupRoutesV1(app fiber.Router) {
	// Setup Configuration
	configuration := config.New()
	database := config.ConnectMongoDatabase(configuration)

	// Setup Repository
	//productRepository := repository.NewProductRepository(database)
	hadithRepository := repositories.NewHadithRepository(database)

	// Setup Service
	//productService := service.NewProductService(&productRepository)
	hadithService := services.NewHadithService(&hadithRepository)

	// Setup Controller
	//productController := controller.NewProductController(&productService)
	hadithController := controllers.NewHadithController(&hadithService)

	//productController.Route(app)
	hadithController.Route(app)
}
