package models

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // JWT expires in 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/home" || r.URL.Path == "/login" || r.URL.Path == "/signup" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther) // Redirect to login page if token is missing
			return
		}

		tokenString := strings.TrimSpace(cookie.Value)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			fmt.Println(claims)
			fmt.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			username := claims.Username
			if r.URL.Path == "/addQty" || r.URL.Path == "/delete" || r.URL.Path == "/acceptRequest" || r.URL.Path == "/declineRequest" || r.URL.Path == "/admin/addBook" || r.URL.Path == "/admin/requests" || r.URL.Path == "/makeAdmin" || r.URL.Path == "/admin/add_book" || r.URL.Path == "/admin/books" {
				err := TypeChecker(username, "admin")
				if err == nil {
					next.ServeHTTP(w, r)
				} else {
					http.Redirect(w, r, "/profile", http.StatusSeeOther)
				}
			} else {
				err := TypeChecker(username, "client")
				if err == nil {
					next.ServeHTTP(w, r)
				} else {
					http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
				}
			}
		}
	})
}
func TypeChecker(username, Usertype string) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	var CorrectUser bool
	err = db.QueryRow(`SELECT EXISTS (select 1 from users where username=? and type=?)`, username, Usertype).Scan(&CorrectUser)
	if err != nil {
		return err
	} else if CorrectUser == false {
		newError := errors.New("Wrong User Type")
		return newError
	} else {
		return nil
	}
}
