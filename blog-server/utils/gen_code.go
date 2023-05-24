package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandCode(size int) string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}
