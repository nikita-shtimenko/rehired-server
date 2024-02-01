package controller

import "github.com/labstack/echo/v4"

func BootstrapRoutes(router *echo.Echo, commonController *commonController, userController *userController) {
	commonController.bootstrapRoutes(router)
	userController.bootstrapRoutes(router)
}

func (cc *commonController) bootstrapRoutes(router *echo.Echo) {
	router.GET("/", cc.homepage)
}

func (uc *userController) bootstrapRoutes(router *echo.Echo) {
	router.GET("/users", uc.GetAll)
}
