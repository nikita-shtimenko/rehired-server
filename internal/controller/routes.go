package controller

import "github.com/labstack/echo/v4"

func BootstrapRoutes(router *echo.Echo, commonController *commonController, userController *userController) {
	commonController.BootstrapRoutes(router)
	userController.BootstrapUserRoutes(router)
}

func (cc *commonController) BootstrapRoutes(router *echo.Echo) {
	router.GET("/", cc.homepage)
}

func (uc *userController) BootstrapUserRoutes(router *echo.Echo) {
	router.GET("/users", uc.GetAll)
}
