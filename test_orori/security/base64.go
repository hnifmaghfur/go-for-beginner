package security

import (
	"encoding/base64"
)

func Base64Encode(s string) (string) {
	data := []byte(s)
	encode := base64.StdEncoding.EncodeToString(data)

	return encode
}
