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

// GetModuleGorm get one module
func (m Module) GetOneModuleGorm(id int) (Module, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&m, id)
	return m, response.Error
}

// UpdateModuleGorm saved module edit
func (m Module) UpdateModuleGorm(id int, mod *Module) (Module, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&m).Where("id = ?", id).Updates(Module{Description: mod.Description,})
	return m, response.Error
}

// DeleteModuleGorm delete client
func (m Module) DeleteModuleGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&m, id)
	return response.Error
}

// ShowModuleGorm show client
func (m Module) ShowModuleGorm() ([]Module, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
  var objects []Module
	response := db.Find(&objects)
	return objects, response.Error
}

// ShowModuleGorm show module
// func (m Module) ShowModuleGorm() ([]Module, error) {
// 	conn := lib.NewConfig()
// 	db := conn.DsnStringGorm()
// 	db.AutoMigrate(&Module{})
// 	rows, err := db.Order("id asc").Model(&m).Rows()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	var response []Module
// 	for rows.Next() {
// 		var item Module
// 		db.ScanRows(rows, &item)
// 		response = append(response, item)
// 	}
// 	return response, err
// }