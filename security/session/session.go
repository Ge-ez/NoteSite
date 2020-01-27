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
