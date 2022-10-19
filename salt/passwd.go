package salt

import (
	"bytes"

	"golang.org/x/crypto/bcrypt"
)

func CreatePasswd(password string) (string, bool) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		encodePW := string(hash)
		return encodePW, true
	} else {
		return "", false
	}
}

func SetPasswd(passwd, random string) string {
	var buffer bytes.Buffer
	buffer.WriteString(passwd)
	buffer.WriteString(random)
	passwdRandom := buffer.String()
	return passwdRandom
}

func CheckPasswd(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}
