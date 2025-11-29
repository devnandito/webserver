package models

import (
	"encoding/json"
	"time"

	"github.com/devnandito/webserver/lib"
	"gorm.io/gorm"
)

type Contribution struct {
	gorm.Model
	Amount int `json:amount`
	Contribution_date time.Time `json:contribution_date`
	Method_pay string `json:method_pay`
	Description string `json:description`
	ClientID int `json:clientid`
	Client Client
	ProductID int `json:productid`
	Product Product
}

func (c *Contribution) ToJson(ct Contribution)([]byte, error) {
	return json.Marshal(ct)
}

func (c *Contribution) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// CreateContributionGorm created a new contribution
func (c Contribution) CreateContributionGorm(ct *Contribution) (Contribution, error){
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Contribution{})
	response := db.Create(&ct)
	data := Contribution {
		Amount: ct.Amount,
		Contribution_date: ct.Contribution_date,
		Method_pay: ct.Method_pay,
		Description: ct.Description,
		ClientID: ct.ClientID,
		ProductID: ct.ProductID,
	}
	
	return data, response.Error
}

// GetContributionGorm get one contribution
func (c Contribution) GetOneContributionGorm(id int) (Contribution, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&c, id)
	return c, response.Error
}

// UpdateContributionGorm saved contribution edit
func (c Contribution) UpdateContributionGorm(id int, ct *Contribution) (Contribution, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&c).Where("id = ?", id).Updates(Contribution{Amount: ct.Amount, Method_pay: ct.Method_pay, Description: ct.Description, ClientID: ct.ClientID, ProductID: ct.ProductID})
	return c, response.Error
}

// DeleteContributionGorm delete contribution
func (c Contribution) DeleteContributionGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&c, id)
	return response.Error
}

// ShowContributionGorm show contribution
func (c Contribution) ShowContributionGorm() ([]Contribution, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
  var objects []Contribution
	response := db.Preload("Client").Preload("Product").Find(&objects)
	return objects, response.Error
}