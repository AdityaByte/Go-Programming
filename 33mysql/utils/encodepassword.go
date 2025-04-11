package utils

import (
	"bytes"
	"encoding/base64"
)

func EncodePassword(pass string) (string, error) {
	// For Encoding the password we needs an encoder.
	var encodedPassword bytes.Buffer
	encoder := base64.NewEncoder(base64.RawStdEncoding.Strict(), &encodedPassword)
	if _, err := encoder.Write([]byte(pass)); err != nil {
		return "", err
	}
	encoder.Close()
	return encodedPassword.String(), nil;
}