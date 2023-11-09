package controllers

import (
	"net/http"
	"os"
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

func generateToken(c *gin.Context, id int, name string, email string, userType bool) (string, error) {
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

	// Set token dalam cookies
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     tokenName,
		Value:    jwtToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})

	return jwtToken, nil
}

func validateTokenFromCookies(r *http.Request) (*Claims, bool) {
	cookie, err := r.Cookie(tokenName)
	if err != nil {
		return nil, false
	}

	jwtToken := cookie.Value
	accessClaims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(jwtToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return accessClaims, err == nil && parsedToken.Valid
}

func validateUserToken(c *gin.Context, accessType bool) bool {
	accessClaims, isValidToken := validateTokenFromCookies(c.Request)
	if !isValidToken {
		return false
	}

	isUserValid := accessClaims.UserType == accessType
	if isUserValid {
		return true
	}

	return false
}

func sendUnAuthorizedResponse(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}

func Authenticate(accessType bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidToken := validateUserToken(c, accessType)
		if !isValidToken {
			sendUnAuthorizedResponse(c)
		} else {
			c.Next()
		}
	}
}
