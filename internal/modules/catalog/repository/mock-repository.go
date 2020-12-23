package repository

import (
	mock "CarCatalog/internal/DB"
	"CarCatalog/internal/modules/catalog/entity"
)

type repository struct{}

func NewMockRepository() ICatalogRepository {
	return &repository{}
}

func (*repository) FindByAllParameters(catalog *entity.Catalog) (*[]entity.Catalog, error){
	return mock.FindByAllParameters(*catalog), nil
}

func (*repository) Create(catalog *entity.Catalog) (*entity.Catalog, error) {
	return mock.CREATE(*catalog), nil
}

func (*repository) FindAll() (*[]entity.Catalog, error) {
	return mock.SELECT(), nil
}

func (*repository) FindCatalogById(id int) (*entity.Catalog, error) {
	return mock.SELECTById(id), nil
}

func (*repository) UpdateCatalog(catalog *entity.Catalog) (*entity.Catalog, error) {
	return mock.UPDATE(*catalog), nil
}

func (*repository) DeleteCatalog(id int) string {
	return mock.DELETE(id)
}