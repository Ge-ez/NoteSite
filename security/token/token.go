package token
import (
	"github.com/dgrijalva/jwt-go"
)
// CustomClaims specifies custom claims
type CustomClaims struct {
	username string `json:"username"`
	jwt.StandardClaims
}
// Generate generates jwt token
func Generate(signingKey []byte, claims jwt.Claims) (string, error) {
	tn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tn.SignedString(signingKey)
	return signedString, err
}
