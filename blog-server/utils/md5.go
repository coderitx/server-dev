package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(byteData []byte) string {
	h := md5.New()
	h.Write(byteData)
	return hex.EncodeToString(h.Sum(nil))
}
