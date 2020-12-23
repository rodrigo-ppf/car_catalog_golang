package repository

import (
	"CarCatalog/internal/modules/catalog/entity"
)

type ICatalogRepository interface {
	Create(catalog *entity.Catalog) (*entity.Catalog, error)
	FindAll() (*[]entity.Catalog, error)
	FindByAllParameters(catalog *entity.Catalog) (*[]entity.Catalog, error)
	FindCatalogById(id int) (*entity.Catalog, error)
	UpdateCatalog(catalog *entity.Catalog) (*entity.Catalog, error)
	DeleteCatalog(id int) string
}
