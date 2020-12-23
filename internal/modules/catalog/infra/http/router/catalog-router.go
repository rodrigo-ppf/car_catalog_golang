package catalogRouter

import (
	controller "CarCatalog/internal/modules/catalog/controller"

	"github.com/gin-gonic/gin"
)

var (
	catalogController controller.ICatalogController = controller.NewCatalogController()
)

func InitializeCatalogRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/catalog/", catalogController.FindAll)
	routerGroup.GET("/catalog/query/", catalogController.FindByAllParameters)
	routerGroup.GET("/catalog/id/:id/", catalogController.FindCatalogById)
	routerGroup.PUT("/catalog/", catalogController.UpdateCatalog)
	routerGroup.DELETE("/catalog/:id/", catalogController.DeleteCatalog)
	routerGroup.POST("/catalog/", catalogController.CreateCatalog)
}
