package utils

import (
	"crypto/md5"
	"fmt"
)

func String2Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}
