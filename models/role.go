package models

import (
	"encoding/json"
	_ "time"

	"github.com/devnandito/webserver/lib"
	
	"gorm.io/gorm"
)

// Role access public
type Role struct {
	// ID uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
	gorm.Model
	Description string `json:"description"`
	Operation []Operation `gorm:"many2many:role_operations;"`
}

// ToJson return to r.body to json
func (r *Role) ToJson(rl Role) ([]byte, error) {
	return json.Marshal(rl)
}

// ToText return r.body to text
func (r *Role) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// ShowRoleGorm show role
func (r Role) ShowRoleGorm() ([]Role, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Role{})
	rows, err := db.Order("id asc").Model(&r).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Role
	for rows.Next() {
		var item Role
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

// CreateRoleGorm insert a new role
func (r Role) CreateRoleGorm(data *Role) (Role, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Role{})
	response := db.Create(&data)
	role := Role {
		Description: r.Description,
	}
	return role, response.Error
}

// UpdateRoleGorm  role edit
func (r Role) UpdateRoleGorm(id int, data *Role) (Role, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&r).Where("id = ?", id).Update("description", data.Description)	
	return r, response.Error
}