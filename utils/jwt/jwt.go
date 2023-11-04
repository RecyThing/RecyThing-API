package jwt

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	godotenv.Load()
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
		TokenLookup:   "cookie:token",
	})
}

func CreateTokenUsers(id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func SetTokenCookie(e echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Path = "/"

	e.SetCookie(cookie)
}

func ExtractTokenUsers(c echo.Context) (string, error) {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok || user == nil {
		return "", errors.New("missing user token")
	}

	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			return "", errors.New("user ID claim is missing or not a string")
		}
		return id, nil
	}

	return "", errors.New("invalid or expired token")
}
