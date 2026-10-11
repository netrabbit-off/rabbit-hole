package helpers

import "encoding/base64"

func DecodeB64(s string) []byte {
	decodedString, _ := base64.StdEncoding.DecodeString(s)
	return decodedString
}

func EncodeB64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
