package services

import (
	"fmt"
	"latest/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService struct {
	secretKey string
	issure    string
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: config.GetConfig().JwtSecretKey,
		issure:    "web-api",
	}
}

type Claim struct {
	Sum string `json:"sum"`
	jwt.StandardClaims
}

func (j *JwtService) GenJwtToken(username string) (string, error) {

	claim := &Claim{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    j.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		return "", err
	}

	return t, nil

}

func (j *JwtService) ValidateToken(token string) bool {

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid Token: %v", token)
		}

		return []byte(j.secretKey), nil
	})

	return err == nil
}
