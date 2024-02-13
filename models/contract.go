package models

import (
	"encoding/json"

	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	Provider    string `json:"provider"`
	Number      string  `json:"number"`
	Year        int  `json:"year"`
	Description string `json:"description"`
	Amount      int  `json:"amount"`
	Fee         int  `json:"fee"`
}

// ToJson return to r.body to json
func (c *Contract) ToJson(co Contract) ([]byte, error) {
	return json.Marshal(co)
}

// ToText return r.body to text
func (c *Contract) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// CreateContractGorm insert a new contract
func (c *Contract) CreateContractGorm(data *Contract) (Contract, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Contract{})
	response := db.Create(&data)
	contract := Contract{
		Provider:    c.Provider,
		Number:      c.Number,
		Year:        c.Year,
		Description: c.Description,
		Amount:      c.Amount,
		Fee:         c.Fee,
	}
	return contract, response.Error
}

// GetContractGorm get one contract
func (c Contract) GetOneContractGorm(id int) (Contract, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Contract{})
	response := db.Find(&c, id)
	return c, response.Error
}

// UpdateContractGorm contract edit
func (c Contract) UpdateContractGorm(id int, data *Contract) (Contract, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&c).Where("id = ?", id).Updates(Contract{Provider: data.Provider, Number: data.Number, Year: data.Year, Description: data.Description, Amount: data.Amount, Fee: data.Fee})
	return c, response.Error
}

// DeleteContractGorm delete contract
func (c Contract) DeleteContractGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&c, id)
	return response.Error
}

// ShowContractGorm show contract
func (c Contract) ShowContractGorm() ([]Contract, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	var objects []Contract
	response := db.Find(&objects)
	return objects, response.Error
}