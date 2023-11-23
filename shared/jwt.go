package shared

import (
	"time"

	"github.com/adityatresnobudi/library-api/dto"
	"github.com/golang-jwt/jwt/v5"
)

var APPLICATION_NAME = "Library"
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("this is a secret")

type JWTClaims struct {
	jwt.RegisteredClaims
	ID int `json:"id"`
}

func AuthorizedJWT(claims JWTClaims, user dto.UserPayload) (string, error) {
	claims.Issuer = APPLICATION_NAME
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims.ID = user.ID

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	generateToken, err := token.SignedString([]byte(JWT_SIGNATURE_KEY))
	if err != nil {
		return "", err
	}

	return generateToken, nil
}

func ValidateJWT(generateToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(generateToken, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return JWT_SIGNATURE_KEY, nil
	})
}
