package models

import (
	"time"

	"github.com/devnandito/webserver/lib"
)

// Role access public
type Role struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Description string `json:"description"`
	Operation []Operation `gorm:"many2many:role_operations;"`
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
	response := db.Create(&data)
	role := Role {
		ID: r.ID,
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