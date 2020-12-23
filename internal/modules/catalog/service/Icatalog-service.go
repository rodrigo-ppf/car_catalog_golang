package service

import (
	"CarCatalog/internal/modules/catalog/entity"
)

type ICatalogService interface {
	Validate(catalog *entity.Catalog) error
	CreateCatalog(catalog *entity.Catalog) (*entity.Catalog, error)
	FindAll() (*[]entity.Catalog, error)
	FindByAllParameters(catalog *entity.Catalog) (*[]entity.Catalog, error)
	FindCatalogById(id int) (*entity.Catalog, error)
	UpdateCatalog(catalog *entity.Catalog) (*entity.Catalog, error)
	DeleteCatalog(id int) string
}
