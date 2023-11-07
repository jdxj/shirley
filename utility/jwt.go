package utility

import (
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.RegisteredClaims
}

func NewToken(key any) (string, error) {
	mc := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "shirley",
			Subject:   "share",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        grand.S(8),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	return token.SignedString(key)
}

func ParseToken(key any, tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (any, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if mc, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return mc, nil
	} else {
		return nil, gerror.NewCode(gcode.CodeValidationFailed)
	}
}
