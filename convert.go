package sid

import (
	"encoding/base64"
	"fmt"
	"log"
)

func encode(bytes []byte) string {
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func decode(base64Str string) (bytes []byte, err error) {
	if base64Str == "" {
		err = fmt.Errorf("base64Str is empty")
		log.Println(err)
		return
	}
	if bytes, err = base64.RawURLEncoding.DecodeString(base64Str); err != nil {
		log.Println(err)
		return
	}
	return
}
