package models

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtKey []byte

type JWTConfig struct {
	JWT_SECRET string `yaml:"JWT_SECRET"`
}

func JwtSecretKey() {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config JWTConfig
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	jwtKey = []byte(config.JWT_SECRET)
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 9).Unix(),
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
		pathComponents := strings.Split(r.URL.Path, "/")
		firstPartOfURL := pathComponents[1]
		if r.URL.Path == "/" || r.URL.Path == "/logout" || firstPartOfURL == "static" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("jwt")
		if err != nil {
			if r.URL.Path == "/" || r.URL.Path == "/home" || r.URL.Path == "/login" || r.URL.Path == "/signup" || r.URL.Path == "/logout" || firstPartOfURL == "static" {
				next.ServeHTTP(w, r)
				return
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}

		tokenString := strings.TrimSpace(cookie.Value)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			fmt.Println(claims)
			fmt.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			username := claims.Username
			if firstPartOfURL == "admin" {
				err := TypeChecker(username, "admin")
				if err == nil {
					if r.URL.Path == "/signup" || r.URL.Path == "/login" || r.URL.Path == "/home" {
						http.Redirect(w, r, "/client/profile", http.StatusSeeOther)
						return
					}
					next.ServeHTTP(w, r)
				} else {
					http.Redirect(w, r, "/client/profile", http.StatusSeeOther)
				}
			} else {
				err := TypeChecker(username, "client")
				if err == nil {
					next.ServeHTTP(w, r)
				} else {
					err = TypeChecker(username, "Requested")
					if err == nil {
						if r.URL.Path == "/signup" || r.URL.Path == "/login" || r.URL.Path == "/home" {
							http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
							return
						}
						next.ServeHTTP(w, r)
					} else {
						http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
					}
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
	err = db.QueryRow(`SELECT EXISTS (SELECT 1 FROM users WHERE username=? and type=?)`, username, Usertype).Scan(&CorrectUser)
	if err != nil {
		return err
	} else if CorrectUser == false {
		newError := errors.New("Wrong User Type")
		return newError
	} else {
		return nil
	}
}
