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
// Valid validates a given token
func Valid(signedToken string, signingKey []byte) (bool, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*CustomClaims); !ok || !token.Valid {
		return false, err
	}

	return true, nil
}
// Claims returns claims used for generating jwt tokens
func Claims(username string, tokenExpires int64) jwt.Claims {
	return CustomClaims{username,
		jwt.StandardClaims{
			ExpiresAt: tokenExpires,
		},
	}
}
