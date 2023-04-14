package test

import (
	"cloud-disk/cloudapi/utils"
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	for i := 0; i < 3; i++ {
		fmt.Println(utils.RandCode())
	}
}
