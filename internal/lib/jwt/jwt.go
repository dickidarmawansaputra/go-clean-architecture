package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type TokenPayload struct {
	ID uint `json:"id"`
}

func Generate(config *viper.Viper, payload *TokenPayload) (string, error) {
	duration, err := time.ParseDuration(config.GetString("JWT_EXP"))
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  payload.ID,
		"exp": time.Now().Add(duration).Unix(),
	})

	token, err := t.SignedString([]byte(config.GetString("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, err

}
