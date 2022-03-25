package data

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

var User = new(usersvc)

type usersvc struct{}

func (usersvc) HashPassword(pwd, salt string) string {
	passwd := pbkdf2.Key([]byte(pwd), []byte(salt), 10000, 50, sha256.New)
	return fmt.Sprintf("%x", passwd)
}
