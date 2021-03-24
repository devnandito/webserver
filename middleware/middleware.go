package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devnandito/webserver/models"
)

func CheckAuth() models.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking authentication")
			if flag {
				f(w, r)
			}else {
				return
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