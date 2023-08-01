package models

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"time"
)

type JWTConfig struct {
	JwtSecretKey string `yaml:"JWT_SECRET"`
}

var config JWTConfig

func init() {
	// Read the configuration from config.yaml
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Error reading config.yaml:", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error parsing config.yaml:", err)
	}
}
func generateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 9).Unix()

	jwtSecret := []byte(config.JwtSecretKey)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
