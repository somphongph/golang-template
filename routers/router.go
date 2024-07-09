package routers

import (
	"golang-template/configs"
	"golang-template/modules/common/repositories/cache"
	"golang-template/modules/examples/handlers"
	"golang-template/modules/examples/repositories"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, cfg *configs.Configs) {
	// API Version
	api := e.Group("/v1")
	cache := cache.InitRedisCache(&cfg.Redis)

	// Stories
	//---------------------------------------------------
	exampleRepo := repositories.InitExampleRepository(&cfg.MongoDB)
	exampleHandler := handlers.ExampleHandler(cfg, exampleRepo, cache)
	exampleApi := api.Group("/stories")
	{
		exampleApi.GET("/:id", exampleHandler.GetItemExampleHandler)
	}

}
