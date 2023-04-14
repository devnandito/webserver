package utils

import (
	"net/http"
	"strings"

	"github.com/devnandito/webserver/models"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)


type ValidateUser struct {
	Email string
	Password string
	Username string
	Name string
	RoleID string
	Errors map[string]string
}

var usr models.User
var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func GetUser(email string) (models.User) {
	res, err := usr.VerifyUser(email)
	if err != nil {
		panic(err)
	}
	data := models.User{
		Email: res.Email,
		Password: res.Password,
		Username: res.Username,
	}
	return data
}

func CompareHash(msgPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword(
 		[]byte(pwd),
 	 	[]byte(msgPwd))
	return err == nil
}

func (msg *ValidateUser) ValidateLogin() bool {
	msg.Errors = make(map[string]string)

	user := GetUser(msg.Email)
	ok := CompareHash(msg.Password, user.Password)

	if !ok {
		msg.Errors["Password"] = "Password incorrect"
	}

	if user.Email != msg.Email {
		msg.Errors["Email"] = "User not exits"
	}

	if strings.TrimSpace(msg.Email) == "" {
		msg.Errors["Email"] = "Please enter a email"
	}

	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please enter a password"
	}
	
	return len(msg.Errors) == 0

}

func (msg *ValidateUser) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Please enter a username"
	}

	if strings.TrimSpace(msg.Name) == "" {
		msg.Errors["Name"] = "Please enter a name"
	}

	if strings.TrimSpace(msg.Email) == "" {
		msg.Errors["Email"] = "Please enter a email"
	}

	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please enter a password"
	}

	if strings.TrimSpace(msg.RoleID) == "" {
		msg.Errors["RoleID"] = "Please enter a rol"
	}
	
	return len(msg.Errors) == 0

}

func (msg *ValidateUser) ValidateEdit() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Please enter a username"
	}

	if strings.TrimSpace(msg.Name) == "" {
		msg.Errors["Name"] = "Please enter a name"
	}

	if strings.TrimSpace(msg.Email) == "" {
		msg.Errors["Email"] = "Please enter a email"
	}

	if strings.TrimSpace(msg.RoleID) == "" {
		msg.Errors["RoleID"] = "Please enter a rol"
	}

	return len(msg.Errors) == 0

}

func GetSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "cookie-name")
	return session
}
