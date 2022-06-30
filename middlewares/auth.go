package middlewares

import (
	"fmt"
	"strings"

	"yas/cfg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt"
)

func JWTAuth(ctx *fiber.Ctx) error {
	store := session.New()
	tokenString := ctx.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer", "", -1)
	tokenString = strings.TrimSpace(tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid Token")
		}
		return cfg.SecretKey(), nil
	})
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		sess, err := store.Get(ctx)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		sess.Set("user", claims.Issuer)
		return ctx.Next()
	}
	return ctx.SendStatus(fiber.StatusUnauthorized)
}
