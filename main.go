package main

import (
	middleware "CarCatalog/internal/shared/infra/helper"
	routers "CarCatalog/internal/shared/infra/http/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	const port string = ":4000"
	router := gin.Default()
	//router.Use(globalExceptionHandler())
	router.Use(middleware.CORS())
	routers.InitializeRoutes(router)
	router.Run(port)

}
