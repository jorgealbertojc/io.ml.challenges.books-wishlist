package signin

import (
	"crypto/rand"
	"encoding/hex"
)

func (l *logic) buildSigninAuthenticationToken() (string, error) {

	length := 60
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
