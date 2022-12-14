package utils

import "golang.org/x/crypto/bcrypt"

// EncodePwd 加密
func EncodePwd(pwd *string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*pwd), bcrypt.DefaultCost)
	*pwd = string(hash)
	return err
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
