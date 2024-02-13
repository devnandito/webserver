package models

import (
	"encoding/json"

	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)

type Period struct {
	gorm.Model
	Start string `json:"start"`
	Finish string `json:"finish"`
}

// ToJson return to r.body to json
func (p *Period) ToJson(pe Period) ([]byte, error) {
	return json.Marshal(pe)
}

// ToText return r.body to text
func (p *Period) ToText (data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// CreatePeriodGorm insert a new period
func (p *Period) CreatePeriodGorm(data *Period) (Period, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Period{})
	response := db.Create(&data)
	period := Period {
		Start: p.Start,
		Finish: p.Finish,
	}
	return period, response.Error
}

// GetOnePeriodGorm get one period
func (p Period) GetOnePeriodGorm(id int) (Period, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Period{})
	response := db.Find(&p, id)
	return p, response.Error
}

// UpdatePeriodGorm period edit
func (p Period) UpdatePeriodGorm(id int, data *Period) (Period, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&p).Where("id = ?", id).Updates(Period{Start: data.Start, Finish: data.Finish})
	return p, response.Error
}

// DeletePeriodGorm delete period
func (p Period) DeletePeriodGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&p, id)
	return response.Error
}

// ShowPeriodGorm show period
func (p Period) ShowPeriodGorm() ([]Period, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	var objects []Period
	response := db.Find(&objects)
	return objects, response.Error
}