package controllers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))
var tokenName = "token"

type Claims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType bool   `json:"userType"`
	jwt.StandardClaims
}

func generateToken(id int, name string, email string, userType bool) (string, error) {
	tokenExpiryTime := time.Now().Add(1000 * 365 * time.Hour)

	// Membuat claims
	claims := &Claims{
		ID:       id,
		Name:     name,
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	// Membuat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func validateTokenFromHeader(r *http.Request) (*Claims, bool) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, false
	}

	// Format header Authorization: Bearer {token}
	splitToken := strings.Split(authorizationHeader, "Bearer ")
	if len(splitToken) != 2 {
		return nil, false
	}

	jwtToken := strings.TrimSpace(splitToken[1])
	accessClaims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(jwtToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return accessClaims, err == nil && parsedToken.Valid
}

func validateUserToken(c *gin.Context, accessType bool) bool {
	accessClaims, isValidToken := validateTokenFromHeader(c.Request)
	if !isValidToken {
		return false
	}

	isUserValid := accessClaims.UserType == accessType
	return isUserValid
}

func sendUnauthorizedResponse(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}

func Authenticate(accessType bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidToken := validateUserToken(c, accessType)
		if !isValidToken {
			sendUnauthorizedResponse(c)
		} else {
			c.Next()
		}
	}
}
