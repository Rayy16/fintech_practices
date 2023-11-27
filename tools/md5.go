package tools

import (
	"crypto/md5"
	"fmt"
)

const Salt = "CC-fintech-practices"

func GenMD5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}

func GenMD5WithSalt(content, salt string) string {
	return GenMD5(content + salt)
}
