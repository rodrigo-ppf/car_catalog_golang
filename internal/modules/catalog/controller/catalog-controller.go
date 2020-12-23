package controller

import (
	"CarCatalog/internal/modules/catalog/entity"
	"CarCatalog/internal/modules/catalog/repository"
	"CarCatalog/internal/modules/catalog/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type controller struct {
}

var (
	catalogService = &service.Service{Repository: repository.NewMockRepository()}
)

func NewCatalogController() *controller {
	return &controller{}
}

func (*controller) FindAll(context *gin.Context) {

	catalogs, err := catalogService.FindAll()
	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	context.JSON(http.StatusOK,catalogs)
}

func (*controller) FindByAllParameters(context *gin.Context) {
	var catalog = entity.Catalog{}

	err := context.Bind(&catalog)
	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	catalogs, err := catalogService.FindByAllParameters(&catalog)
	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	out, err := json.Marshal(catalogs)
	if err != nil {
		panic(err)
	}

	context.String(http.StatusOK, string(out))
}

func (*controller) FindCatalogById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	catalog, err := catalogService.FindCatalogById(id)
	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	out, err := json.Marshal(catalog)
	if err != nil {
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	context.String(http.StatusOK, string(out))
}

func (*controller) UpdateCatalog(context *gin.Context) {
	var catalog = entity.Catalog{}

	err := context.Bind(&catalog)
	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	_err := catalogService.Validate(&catalog)
	if _err != nil{
		context.AbortWithStatusJSON(http.StatusInternalServerError,_err.Error())
		return
	}

	car, err := catalogService.UpdateCatalog(&catalog)
	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	out, err := json.Marshal(car)
	if err != nil {
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	context.String(http.StatusOK, string(out))

}

func (*controller) CreateCatalog(context *gin.Context) {
	var catalog = entity.Catalog{}
	_ = context.Bind(&catalog)

	err := catalogService.Validate(&catalog)

	if err != nil{
		context.AbortWithStatusJSON(http.StatusInternalServerError,err.Error())
		return
	}

	car, err := catalogService.CreateCatalog(&catalog)
	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	out, err := json.Marshal(car)
	if err != nil {
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	context.String(http.StatusOK, string(out))
}

func (*controller) DeleteCatalog(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	context.String(http.StatusOK, catalogService.DeleteCatalog(id))
}
