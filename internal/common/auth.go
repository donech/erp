package common

import (
	"encoding/base64"
	"flag"

	"github.com/donech/tool/cipher"
)

var authKey = "donech"

func init() {
	flag.StringVar(&authKey, "authKey", "donech0123456789", "key for encrypt password")
}

func ValidatePassword(password, encryptPassword string) bool {
	return EncryptPassword(password) == encryptPassword
}

func EncryptPassword(password string) string {
	res := cipher.AesEncryptCBC([]byte(password), []byte(authKey))
	return base64.StdEncoding.EncodeToString(res)
}
