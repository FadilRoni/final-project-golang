package helper

import "golang.org/x/crypto/bcrypt"

func HassPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {

	err := bcrypt.CompareHashAndPassword(h, p)

	return err == nil
}
