package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword( hashPwd string, password string ) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
	return err == nil
}
