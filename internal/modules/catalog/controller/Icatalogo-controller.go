package controller

import (
	"github.com/gin-gonic/gin"
)

type ICatalogController interface {
	FindAll(context *gin.Context)
	FindByAllParameters(context *gin.Context)
	FindCatalogById(context *gin.Context)
	UpdateCatalog(context *gin.Context)
	CreateCatalog(context *gin.Context)
	DeleteCatalog(context *gin.Context)
}
