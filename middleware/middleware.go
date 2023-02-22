package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devnandito/webserver/models"
)

var cls models.Client

func CheckAuth() models.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			response, err := cls.GetOneClientGorm(1)

			if err != nil {
				panic(err)
			}

			if response.Ci == "3796986" {
				flag = true
				fmt.Println("User authenticated", response)
				fmt.Println("Checking authentication")
			}

			if flag {
				f(w, r)
			}else {
				fmt.Println("User not authenticated", response)
				http.Redirect(w, r, "/", http.StatusFound)
			}
		}
	}
}

func Logging() models.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}