package token
import (
	"github.com/dgrijalva/jwt-go"
)
// CustomClaims specifies custom claims
type CustomClaims struct {
	username string `json:"username"`
	jwt.StandardClaims
}
