package models

import (
	"encoding/json"

	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)

// Operation access public
type Operation struct {
	gorm.Model
	Description string `json:"description"`
	ModuleID int `json:"moduleid"`
	Module Module
}

// ToJson return to r.body to json
func (o Operation) ToJson(op Operation) ([]byte, error) {
	return json.Marshal(op)
}

// ToText return r.body to text
func (o *Operation) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// CreateOperationGorm created a new operation
func (o Operation) CreateOperationGorm(op *Operation) (Operation, error){
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Operation{})
	response := db.Create(&op)
	data := Operation {
		Description: op.Description,
		ModuleID: op.ModuleID,
	}
	
	return data, response.Error
}

// GetOperationGorm get one module
func (o Operation) GetOneOperationGorm(id int) (Operation, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&o, id)
	return o, response.Error
}

// UpdateOperationGorm saved operation edit
func (o Operation) UpdateOperationGorm(id int, op *Operation) (Operation, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&o).Where("id = ?", id).Updates(Operation{Description: op.Description, ModuleID: op.ModuleID})
	return o, response.Error
}

// DeleteOperationGorm delete operation
func (o Operation) DeleteOperationGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&o, id)
	return response.Error
}

// ShowOperationGorm show operation
func (o Operation) ShowOperationGorm() ([]Operation, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
  var objects []Operation
	response := db.Preload("Module").Find(&objects)
	return objects, response.Error
}

// ShowOperationGormPreload
// func (o Operation) ShowOperationGormPreload() ([]Operation, error) {
// 	conn := lib.NewConfig()
// 	db := conn.DsnStringGorm()
// 	var objects []Operation
// 	response := db.Preload("Module").Find(&objects)
// 	return objects, response.Error
// }