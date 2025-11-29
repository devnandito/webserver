package models

import (
	"encoding/json"

	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	Description string `json:description`
}

func (p *Product) ToJson(ct Product)([]byte, error) {
	return json.Marshal(ct)
}

func (p *Product) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// CreateProductGorm created a new product
func (p Product) CreateProductGorm(pr *Product) (Product, error){
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Create(&pr)
	data := Product {
		Description: pr.Description,
	}
	
	return data, response.Error
}

// GetProductGorm get one Product
func (p Product) GetOneProductGorm(id int) (Product, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&p, id)
	return p, response.Error
}

// UpdateProductGorm saved Product edit
func (p Product) UpdateProductGorm(id int, pr *Product) (Product, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&p).Where("id = ?", id).Updates(Product{Description: pr.Description})
	return p, response.Error
}

// DeleteProductGorm delete Product
func (p Product) DeleteProductGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&p, id)
	return response.Error
}

// ShowProductGorm show Product
func (c Product) ShowProductGorm() ([]Product, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Product{})
  var objects []Product
	response := db.Find(&objects)
	return objects, response.Error
}