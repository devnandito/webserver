package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/devnandito/webserver/lib"

	"gorm.io/gorm"
)

// Client client access public
type Client struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Ci string `json:"ci"`
	Birthday time.Time `json:"birthday"`
	Sex string `json:"sex"`
	Nationality string `json:"nationality"`
	DesType string `json:"destype"`
	Code1 string `json:"code1"`
	Code2 string `json:"code2"`
	Code3 string `json:"code3"`
	Direction string `json:"direction"`
	Phone string `json:"phone"`
}

// BirthdayDateStr conver to string
func (c Client) BirthdayDateStr() string {
	return c.Birthday.Format("2006-01-02")
}
// BirthdayTime convert string to time
func (c Client) BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}

// ToJson return to r.body to json
func (c *Client) ToJson(cls Client) ([]byte, error) {
	return json.Marshal(cls)
}

// ToText return r.body to text
func (c *Client) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// FormClient form access public
type FormClient struct {
	Ci string
	FirstName string
	LastName string
}

// GetClientGorm show all client
func (c Client) GetClientGorm(fcls *FormClient) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	name := strings.ToUpper(fcls.FirstName)
	last := strings.ToUpper(fcls.LastName)
	// name := strings.Title(fcls.FirstName)
	// last := strings.Title(fcls.LastName)
	var condition string
	var value1 string
	var value2 string
	var value3 string
	var val []string
	if name == "" && last == "" {
		condition = "ci = ?"
		val = append(val, fcls.Ci)
	} else if fcls.Ci == "" && last == "" {
		condition = "first_name LIKE ? "
		val = append(val, name+"%")
	} else if fcls.Ci == "" && name == "" {
		condition = "last_name LIKE ?"
		val = append(val, last+"%")
	} else if fcls.Ci == "" {
		condition = "last_name LIKE ? OR first_name LIKE ?"
		val = append(val, last+"%")
		val = append(val, name+"%")
		} else if name == "" {
			condition = "last_name LIKE ? OR ci = ?"
			val = append(val, last+"%")
			val = append(val, fcls.Ci)
	} else if last == "" {
		condition = "first_name LIKE ? OR ci = ?"
		val = append(val, name+"%")
		val = append(val, fcls.Ci)
	} else {
		condition = "first_name LIKE ? OR last_name LIKE ? OR ci = ?"
		val = append(val, name+"%")
		val = append(val, last+"%")
		val = append(val, fcls.Ci)
	}
	if len(val) == 3 {
		value1 = val[0]
		value2 = val[1]
		value3 = val[2]
	} else if len(val) == 2 {
		value1 = val[0]
		value2 = val[1]
	} else {
		value1 = val[0]
		value2 = ""
		value3 = ""
	}
	rows, err := db.Order("id asc").Model(&c).Where(condition, value1, value2, value3).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Client
	for rows.Next() {
		var item Client
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

// ApiGetClientGorm show all client
func (c Client) ApiGetClientGorm(ci string) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	rows, err := db.Order("id asc").Model(&c).Where("ci = ?", ci).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []Client
	for rows.Next() {
		var item Client
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

//CreateClientGorm insert new client
func (c Client) CreateClientGorm(cls *Client) (Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Client{})
	response := db.Create(&cls)
	data := Client{
		FirstName: cls.FirstName,
		LastName: cls.LastName,
		Ci: cls.Ci,
		Birthday: cls.Birthday,
		Sex: cls.Sex,
	}
	return data, response.Error
}

// GetClientGorm get one client
func (c Client) GetOneClientGorm(id int64) (Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&c, id)
	return c, response.Error
}

// UpdateClientGorm saved client edit
func (c Client) UpdateClientGorm(id int, cls *Client) (Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&c).Where("id = ?", id).Updates(Client{FirstName: cls.FirstName, LastName: cls.LastName, Ci: cls.Ci, Birthday: cls.Birthday, Sex: cls.Sex})
	return c, response.Error
}

// DeleteClientGorm delete client
func (c Client) DeleteClientGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&c, id)
	return response.Error
}

// ShowClientGorm show client
func (c Client) ShowClientGorm() ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
  var clients []Client
	response := db.Find(&clients)
	return clients, response.Error
}

// ShowClientGorm show client
// func (c Client) ShowClientGorm() ([]Client, error) {
// 	conn := lib.NewConfig()
// 	db := conn.DsnStringGorm()
// 	db.AutoMigrate(&Client{})
// 	rows, err := db.Order("id asc").Model(&c).Rows()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	var response []Client
// 	for rows.Next() {
// 		var item Client
// 		db.ScanRows(rows, &item)
// 		response = append(response, item)
// 	}
// 	return response, err
// }
