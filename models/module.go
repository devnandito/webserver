package models

import (
	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)

// Module access public
type Module struct {
	gorm.Model
	Description string `json:"description"`
}

// ShowModuleGorm show module
func (m Module) ShowModuleGorm() ([]Module, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Module{})
	rows, err := db.Order("id asc").Model(&m).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Module
	for rows.Next() {
		var item Module
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}