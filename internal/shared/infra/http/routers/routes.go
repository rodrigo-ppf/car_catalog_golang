package router

import (
	"github.com/gin-gonic/gin"
	catalogRouter "CarCatalog/internal/modules/catalog/infra/http/router"
)

func InitializeRoutes(httpRouter *gin.Engine) {
	catalog := httpRouter.Group("/v1/")
	catalogRouter.InitializeCatalogRoute(catalog)
}
