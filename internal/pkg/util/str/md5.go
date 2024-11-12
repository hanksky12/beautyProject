package str

import (
	"crypto/md5"
	"fmt"
	"io"
)

func HashPassword(password string) string {
	salt := "@#$%"
	h := md5.New()
	io.WriteString(h, password+salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}
