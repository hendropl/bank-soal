package helper

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func Write(c *gin.Context, token string) error {
	if token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	c.SetCookie(
		"token",
		token,
		3600*24*7,
		"/",
		"",
		false,
		true,
	)

	return nil
}

func Read(c *gin.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		return fmt.Errorf("cookie not found: %w", err)
	}

	if token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	return nil
}
