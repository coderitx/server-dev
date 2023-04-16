package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
