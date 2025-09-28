package helper

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"latih.in-be/model"
)

type ClaimsModel struct {
	UserID int    `json:"id"`
	Role   string `json:"role"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func ParseExpiry(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("empty expiry string")
	}

	if strings.HasSuffix(s, "d") {
		nStr := strings.TrimSuffix(s, "d")
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return 0, err
		}
		return time.Duration(n) * 24 * time.Hour, nil
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}
	return d, nil
}

func GenerateToken(user *model.User, expiryStr string) (string, time.Time, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", time.Time{}, errors.New("JWT_SECRET is not set in environment")
	}

	if expiryStr == "" {
		expiryStr = os.Getenv("JWT_EXPIRED")
	}
	if expiryStr == "" {
		expiryStr = "7d"
	}

	duration, err := ParseExpiry(expiryStr)
	if err != nil {
		return "", time.Time{}, err
	}

	expireAt := time.Now().Add(duration)

	claims := ClaimsModel{
		UserID: user.Id,
		Role:   string(user.Role),
		Name:   user.Name,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireAt),
			Subject:   strconv.Itoa(user.Id),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return ss, expireAt, nil
}

func ParseAndValidateToken(tokenStr string) (*ClaimsModel, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET is not set in environment")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &ClaimsModel{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimsModel)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
