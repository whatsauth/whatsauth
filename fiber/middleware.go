package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
)

type AuthMiddlewareFiber struct {
	PublicKey   string
	PrivateKey  string
	AuthHeader  string
	TokenHeader string
}

func AuthMiddleware(publicKey string, privateKey string, authHeader string, tokenHeader string) *AuthMiddlewareFiber {
	return &AuthMiddlewareFiber{
		PublicKey:   publicKey,
		PrivateKey:  privateKey,
		AuthHeader:  authHeader,
		TokenHeader: tokenHeader,
	}
}

func (auth AuthMiddlewareFiber) GetAndDecodeToken(ctx *fiber.Ctx) (err error) {
	tokenString := ctx.Get(auth.AuthHeader)
	if tokenString == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Header Not Found")
	}

	stringDec, err := watoken.Decode(auth.PublicKey, tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unknown Format Header")
	}

	ctx.Set(auth.TokenHeader, stringDec.Id)
	err = ctx.Next()
	return
}
