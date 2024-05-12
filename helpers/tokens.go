package helpers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("keyrahasiasuper")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token")
	} else if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ProtectedHandler(c *gin.Context) bool {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		log.Println("Missing authorization header")
		return false
	}

	if strings.HasPrefix(tokenString, "Bearer") {
		tokenString = tokenString[len("Bearer "):]
	}

	if err := VerifyToken(tokenString); err != nil {
		log.Println("Invalid token")
		return false
	}
	return true
}

func GetusernameFromToken(c *gin.Context) string {
	username := ""
	tokenString := c.GetHeader("Authorization")
	if strings.HasPrefix(tokenString, "Bearer") {
		tokenString = tokenString[len("Bearer "):]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Println(err.Error())
		return username
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = fmt.Sprint(claims["username"])
	}

	return username
}
