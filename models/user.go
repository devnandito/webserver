package models

import (
	"encoding/json"

	"github.com/devnandito/webserver/lib"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// type User struct {
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Phone string `json:"phone"`
// }

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	RoleID int `json:"role"`
	Role Role
}

func (u User) GetPwdHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
 	return string(hash), err
}

// ToJson return to r.body to json
func (u *User) ToJson(usr User) ([]byte, error) {
	return json.Marshal(usr)
}

// ToText return r.body to text
func (u *User) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// CreateUserGorm insert a new user
func (u User) CreateUserGorm(data *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&User{})
	response := db.Create(&data)
	user := User {
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
		RoleID: u.RoleID,
	}

	return user, response.Error
}

// UpdateUserGorm update user
func (u User) UpdateUserGorm(id int, usr *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&u).Where("id = ?", id).Updates(User{Username: usr.Username,Email: usr.Email, Name: usr.Name, RoleID: usr.RoleID})
	return u, response.Error
}

// GetUserGorm get one user
func (u User) GetOneUserGorm(id int) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&u, id)
	return u, response.Error
}

// DeleteDeleteGorm delete user
func (u User) DeleteUserGorm(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&u, id)
	return response.Error
}

// ShowClientGorm show client
func (u User) ShowUserGorm() ([]User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
  var objects []User
	response := db.Find(&objects)
	return objects, response.Error
}

// ChangePassword saved user edit
func (u User) SaveEditUserGorm(id int, usr *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&u).Where("id = ?", id).Updates(User{Username: usr.Username, Email: usr.Email, Password: usr.Password, RoleID: usr.RoleID})
	return u, response.Error
}

// ShowUserGorm show user
func (u User) ShowUserGorm1() ([]User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&User{})
	rows, err := db.Order("id asc").Model(&u).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var response []User
	for rows.Next() {
		var item User
		db.ScanRows(rows, &item)
		response = append(response, item)
	}
	return response, err
}

// Search user
func (u User) SearchUser(data *User) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Where("username = ?", data.Username).Find(&u)
	return u, response.Error
}

// Search user for ID
func (u User) SearchUserID(data string) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Where("id = ?", data).Find(&u)
	return u, response.Error
}

// Verify user
func (u User) VerifyUser(email string) (User, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Where("email = ?", email).Find(&u)
	return u, response.Error
}

// type AuthUser struct {
// 	Email string
// 	PasswordHash string
// }

// var authUserDB = map[string]AuthUser{}

// type UserService struct {
// }

// func (UserService) VerifyUser(user User) bool {
// 	authUser, ok := authUserDB[user.Email]
// 	if !ok {
// 		return false
// 	}
	
// 	err := bcrypt.CompareHashAndPassword(
// 		[]byte(authUser.PasswordHash),
// 		[]byte(user.Password))
// 		return err == nil
// }


// func (u User) VerifyUser(data *User) (bool, error) {
// 	conn := lib.NewConfig()
// 	db := conn.DsnStringGorm()
// 	response := db.Where("email = ?", data.Email).Find(&u)

// 	return true, response.Error

// 	// err := bcrypt.CompareHashAndPassword(
// 	// 	[]byte(data.Password),
// 	// 	[]byte(u.Password))
	
// }

// func (UserService) CreateUser(newUser User) error {
// 	_, ok := authUserDB[newUser.Email]
// 	if !ok {
// 		return errors.New("User already exists")
// 	}

// 	passwordHash, err := GetPasswordHash(newUser.Password)
// 	if err != nil {
// 		return err
// 	}

// 	newAuthUser := AuthUser {
// 		Email: newUser.Email,
// 		PasswordHash: passwordHash,
// 	}

// 	authUserDB[newAuthUser.Email] = newAuthUser
// 	return nil
// }

// func GetPasswordHash(password string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
// 	return string(hash), err
// }
