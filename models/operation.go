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

// ShowOperationGorm show user
func (o Operation) ShowOperationGorm() ([]Operation, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Operation{})
	rows, err := db.Order("id asc").Model(&o).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Operation
	for rows.Next() {
		var item Operation
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
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