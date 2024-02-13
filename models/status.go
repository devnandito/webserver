package models

import (
	"encoding/json"

	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Description string `json:"description"`
}

// ToJson return to r.body to json
func (s *Status) ToJson(st Status) ([]byte, error) {
	return json.Marshal(st)
}

// ToText return r.body to text
func(s *Status) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// CreateStatusGorm insert a new stutus
func (s *Status) CreateStatusGorm(data *Status) (Status, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Status{})
	response := db.Create(&data)
	status := Status{
		Description: s.Description,
	}
	return status, response.Error
}

// GetStatusGorm get one status
func (s Status) GetOneStatusGorm(id int) (Status, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Status{})
	response := db.Find(&s, id)
	return s, response.Error
}

// UpdateStatusGorm status edit
func (s Status) UpdateStatusGorm(id int, data *Status) (Status, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&s).Where("id = ?", id).Updates(Status{Description: data.Description})
	return s, response.Error
}

// DeleteStatusGorm delete status
func (s Status) DeleteStatusGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&s, id)
	return response.Error
}

// ShowStatusGorm show status
func (s Status) ShowStatusGorm() ([]Status, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	var objects []Status
	response := db.Find(&objects)
	return objects, response.Error
}