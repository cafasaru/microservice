package auth

import (
	"github.com/labstack/echo/v4"
)

// HTTPHandler interface is the interface for interacting with the HTTP transport layer
type HTTPHandler interface {
	Login() echo.HandlerFunc
	RefreshToken() echo.HandlerFunc
	VerifyToken() echo.HandlerFunc
}
