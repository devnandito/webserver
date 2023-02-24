package models

import (
	"encoding/json"
	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)

// Module access public
type Module struct {
	gorm.Model
	Description string `json:"description"`
}

// ToJson return to r.body to json
func (m *Module) ToJson(md Module) ([]byte, error) {
	return json.Marshal(md)
}

// ToText return r.body to text
func (m *Module) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
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

//CreateModuleGorm created a new module
func (m Module) CreateModuleGorm(md *Module) (Module, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Module{})
	response := db.Create(&md)
	data := Module{
		Description: md.Description,
	}

	return data, response.Error
}