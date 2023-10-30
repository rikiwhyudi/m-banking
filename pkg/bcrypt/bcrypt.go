package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {

	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedByte), err
}

func CheckPassword(password, hashedPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
