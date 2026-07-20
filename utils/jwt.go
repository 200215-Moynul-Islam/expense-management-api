package utils

import (
	"time"

	appErrors "expense-management-api/errors"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret         []byte
	jwtExpirationTime time.Duration
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func init() {
	secret, err := beego.AppConfig.String("JWT_SECRET")
	if err != nil {
		panic("JWT_SECRET is not configured")
	}

	expirationMinutes, err := beego.AppConfig.Int("JWT_EXPIRATION_MINUTES")
	if err != nil {
		panic("JWT_EXPIRATION_MINUTES is not configured")
	}

	jwtSecret = []byte(secret)
	jwtExpirationTime = time.Duration(expirationMinutes) * time.Minute
}

// GenerateToken generates a JWT token.
func GenerateToken(
	userID int,
) (string, error) {

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(jwtExpirationTime),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(jwtSecret)
}

// ValidateToken validates a JWT token and returns the user ID.
func ValidateToken(
	tokenString string,
) (int, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			return jwtSecret, nil
		},
	)

	if err != nil {
		return 0, appErrors.ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, appErrors.ErrInvalidToken
	}

	return claims.UserID, nil
}
