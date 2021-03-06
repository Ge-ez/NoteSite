pacakge session

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"github.com/notesite/security/token"
	"github.com/dgrijalva/jwt-go"
)
func Create(claims jwt.Claims, sessionID string, signingKey []byte, w http.ResponseWriter) {

	signedString, err := token.Generate(signingKey, claims)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	c := http.Cookie{
		Name:     sessionID,
		Value:    signedString,
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}
// Valid validates client cookie value
func Valid(cookieValue string, signingKey []byte) (bool, error) {
	valid, err := token.Valid(cookieValue, signingKey)
	if err != nil || !valid {
		return false, errors.New("Invalid Session Cookie")
	}
	return true, nil
}

// Remove expires existing session
func Remove(sessionID string, w http.ResponseWriter) {
	c := http.Cookie{
		Name:    sessionID,
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
		Value:   "",
	}
	http.SetCookie(w, &c)
}
