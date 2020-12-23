package mock

import (
	"CarCatalog/internal/modules/catalog/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	err error
)

func validIsEmptyString(valueTest string, replaceWith string) string {
	if valueTest == "" {
		return replaceWith
	}
	return valueTest
}

func validIsEmptyInt(valueTest int, replaceWith int) int {
	if valueTest <= 0 {
		return replaceWith
	}
	return valueTest
}

func validIsEmptyFloat(valueTest float64, replaceWith float64) float64 {
	if valueTest == 0 {
		return replaceWith
	}
	return valueTest
}

func readFile() []entity.Catalog {
	// ready file transformar em função
	file, _ := ioutil.ReadFile("cars.json")
	catalogs := []entity.Catalog{}
	_ = json.Unmarshal([]byte(file), &catalogs)
	return catalogs
}

func nextId(catalogs []entity.Catalog) int {
	max := 0
	for i := 0; i < len(catalogs); i++ {
		if catalogs[i].Id > max {
			max = catalogs[i].Id
		}
	}
	return max + 1
}

func commit(catalogs []entity.Catalog) {
	_file, _ := json.MarshalIndent(catalogs, "", " ")
	_ = ioutil.WriteFile("cars.json", _file, 0644)
}

func SELECT() *[]entity.Catalog {

	var catalogs = readFile()

	return &catalogs
}

func FindByAllParameters(catalog entity.Catalog) *[]entity.Catalog {

	var catalogs = readFile()
	catalogsFilters := []entity.Catalog{}

	if catalog.Brand == "" && catalog.Fuel == "" && catalog.Model == "" && catalog.Year == 0 && catalog.Price == 0 {
		return SELECT()
	}

	for i := 0; i < len(catalogs); i++ {
		if catalogs[i].Brand == validIsEmptyString(catalog.Brand, catalogs[i].Brand) &&
			catalogs[i].Fuel == validIsEmptyString(catalog.Fuel, catalogs[i].Fuel) &&
			catalogs[i].Year == validIsEmptyInt(catalog.Year, catalogs[i].Year) &&
			catalogs[i].Price == validIsEmptyFloat(catalog.Price, catalogs[i].Price) &&
			catalogs[i].Model == validIsEmptyString(catalog.Model, catalogs[i].Model) {
			car := entity.Catalog{
				Id:    catalogs[i].Id,
				Brand: catalogs[i].Brand,
				Model: catalogs[i].Model,
				Year:  catalogs[i].Year,
				Fuel:  catalogs[i].Fuel,
				Price: catalogs[i].Price}
			catalogsFilters = append(catalogsFilters, car)
		}
	}

	return &catalogsFilters

}

func SELECTById(id int) *entity.Catalog {

	var catalogs = readFile()
	var car = entity.Catalog{}

	catalogsFilters := []entity.Catalog{}
	for i := 0; i < len(catalogs); i++ {
		if catalogs[i].Id == id {
			fmt.Println("Brand: " + catalogs[i].Brand)
			car = catalogs[i]
			catalogsFilters = append(catalogsFilters, car)

		}
	}
   	return &car
}

func CREATE(catalog entity.Catalog) *entity.Catalog {

	var catalogs = readFile()
	catalog.Id = nextId(catalogs)

	catalogs = append(catalogs, catalog)

	 commit(catalogs)

	return &catalog
}

func UPDATE(catalog entity.Catalog) *entity.Catalog {

	var catalogs = readFile()

	for i := 0; i < len(catalogs); i++ {

		if catalogs[i].Id == catalog.Id {
			catalogs[i].Id = catalog.Id
			catalogs[i].Brand = catalog.Brand
			catalogs[i].Model = catalog.Model
			catalogs[i].Year = catalog.Year
			catalogs[i].Fuel = catalog.Fuel
			catalogs[i].Price = catalog.Price
		}
	}

	//data := []entity.Catalog{}
	 commit(catalogs)

	return &catalog
}

func DELETE(id int) string {

	var catalogs = readFile()
	var catalog = SELECTById(id)

	catalogsFilters := []entity.Catalog{}

	for i := 0; i < len(catalogs); i++ {

		if catalogs[i].Id != catalog.Id {
			car := entity.Catalog{
				Id:    catalogs[i].Id,
				Brand: catalogs[i].Brand,
				Model: catalogs[i].Model,
				Year:  catalogs[i].Year,
				Fuel:  catalogs[i].Fuel,
				Price: catalogs[i].Price}

			catalogsFilters = append(catalogsFilters, car)
		}
	}

	_file, _ := json.MarshalIndent(catalogsFilters, "", " ")
	_ = ioutil.WriteFile("cars.json", _file, 0644)

	return string("Removido com sucesso!")
}
