package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))
var tokenName = "token"


type Claims struct {
		ID    int    `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
		UserType bool   `json:"userType"`
		jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, name string, email string, userType bool){
	tokenExpiryTime := time.Now().Add(1000*365* time.Hour)

	//create claims
	claims := &Claims{
		ID: 				id,
		Name: 				name,
		Email: 				email,
		UserType: 			userType,
		StandardClaims: 	jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	//encrypt claim to jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtKey)
	if err != nil{
		return
	}

	//set token to cookies
	http.SetCookie(w, &http.Cookie{
		Name:		tokenName,
		Value: 		jwtToken,
		Expires:	tokenExpiryTime,
		Secure: 	false,
		HttpOnly: 	true,
	})
}

func resetUserToken(w http.ResponseWriter){
	http.SetCookie(w, &http.Cookie{
		Name:		tokenName,
		Value: 		"",
		Expires:	time.Now(),
		Secure: 	false,
		HttpOnly: 	true,
	})
}

func validateTokenFromCookies(r *http.Request)(bool, int, string, string, bool){
	if cookie, err := r.Cookie(tokenName); err == nil {
		jwtToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(jwtToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error){
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid{
			return true, accessClaims.ID, accessClaims.Name, accessClaims.Email, accessClaims.UserType
		}
	}
	return false, -1, "", "", false
}

func validateUserToken(r *http.Request, accessType bool) bool {
	isAccessTokenValid, id, name, email, userType :=
	validateTokenFromCookies(r)
	fmt.Print(id, name, email, userType, accessType, isAccessTokenValid)

	if isAccessTokenValid{
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}

	return false
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	panic("unimplemented")
}

func Authenticate(next http.HandlerFunc, accesType bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(r, accesType)
		if !isValidToken {
			sendUnAuthorizedResponse(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}