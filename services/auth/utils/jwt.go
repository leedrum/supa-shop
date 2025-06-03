package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("your-secret-key")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func GenerateJWT(username string, role string, duration time.Duration) (string, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString(jwtKey)
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	if payload.ExpiredAt.IsZero() {
		return nil, fmt.Errorf("missing expired at: %v", payload.ExpiredAt)
	}

	return &jwt.NumericDate{
		Time: payload.ExpiredAt,
	}, nil
}

func (payload *Payload) GetID() (string, error) {
	return payload.ID.String(), nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	if payload.IssuedAt.IsZero() {
		return nil, fmt.Errorf("missing issued at: %v", payload.IssuedAt)
	}

	return &jwt.NumericDate{
		Time: payload.IssuedAt,
	}, nil
}

func (payload *Payload) GetIssuer() (string, error) {
	return payload.Username, nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: time.Now(),
	}, nil
}

func (payload *Payload) GetSubject() (string, error) {
	return payload.Username, nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return jwt.ErrTokenExpired
	}

	return nil
}
