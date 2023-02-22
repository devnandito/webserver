package models

import (
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