package common

import (
	"context"
	"flag"
	"github.com/dgrijalva/jwt-go"
	"github.com/donech/tool/xlog"
	"time"
)

var JWTKey string
var JWTExp string

func init()  {
	flag.StringVar(&JWTKey, "jwt.key", "donech2021", "-jwt.key=demo")
	flag.StringVar(&JWTExp, "jwt.exp", "60m", "-jwt.exp=60m")
}

type CustomField struct {
	UserID int64
	Name string
}

type Claims struct {
	CustomField
	jwt.StandardClaims
}

func GenToken(ctx context.Context, field CustomField) (string, error) {
	duration, err := time.ParseDuration(JWTExp)
	if err != nil {
		xlog.S(ctx).Error("GenToken.ParseDuration error, err=", err)
		return "", err
	}
	now := time.Now()
	claims := Claims{
		CustomField: field,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(duration).Unix(),
			NotBefore: now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		xlog.S(ctx).Error("GenToken.SignedString error, err=", err)
		return "", err
	}
	return ss, nil
}

func RefreshToken(ctx context.Context, token string) (string, error) {
	claims, err := ValidToken(ctx, token)
	if err != nil {
		return "", err
	}
	return GenToken(ctx, claims.CustomField)
}

func ValidToken(ctx context.Context, token string) (Claims, error) {
	claims := Claims{}
	_, err := jwt.ParseWithClaims(token,&claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	return claims, err
}
