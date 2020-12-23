package service

import (
	"CarCatalog/internal/modules/catalog/entity"
	"CarCatalog/internal/modules/catalog/repository"
	"errors"
)

type Service struct {
	Repository repository.ICatalogRepository
}

func NewCatalogService() ICatalogService {
	return &Service{}
}

func (service *Service) FindByAllParameters(catalog *entity.Catalog) (*[]entity.Catalog, error) {
	catalogs, err := service.Repository.FindByAllParameters(catalog)
	if err != nil {
		return nil, err
	}
	return catalogs, nil
}

func (service *Service) Validate(catalog *entity.Catalog) error {
	if catalog.Brand == ""	{
		err := errors.New("Brand is required!")
		return err
	}
	if catalog.Model == ""	{
		err := errors.New("Model is required!")
		return err
	}
	if catalog.Fuel == "" {
		err := errors.New("Fuel is required!")
		return err
	}
	if catalog.Year == 0 {
		err := errors.New("Year is required!")
		return err
	}
	if catalog.Price == 0 {
		err := errors.New("Price is required!")
		return err
	}

	return nil
}

func (service *Service) CreateCatalog(catalog *entity.Catalog) (*entity.Catalog, error) {

	catalog, err := service.Repository.Create(catalog)

	if err != nil{
		return nil, err
	}

	return catalog, nil
}

func (service *Service) FindAll() (*[]entity.Catalog, error)  {

	catalogs, err := service.Repository.FindAll()
	if err != nil{
		return nil, err
	}
	return catalogs, nil

}

func (service *Service) FindCatalogById(id int) (*entity.Catalog, error)  {

	catalog, err := service.Repository.FindCatalogById(id)
	if err != nil{
		return nil, err
	}
	return catalog, nil
}

func (service *Service) UpdateCatalog(catalog *entity.Catalog) (*entity.Catalog, error) {
	catalog, err := service.Repository.UpdateCatalog(catalog)
	if err != nil{
		return nil, err
	}
	return catalog, nil
}

func (service *Service) DeleteCatalog(id int) string {
	return service.Repository.DeleteCatalog(id)
}
