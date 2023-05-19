package tencent

import (
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"testing"
)

func TestUploadTencent(t *testing.T) {
	godotenv.Load("/Users/scliang/Developer/golang/src/server-dev/blog-server/private.env")
	f, err := os.Open("/Users/scliang/Developer/golang/src/server-dev/blog-server/testdata/唯美少女.jpeg")
	imgByte, _ := ioutil.ReadAll(f)
	_, err = UploadImageTencent(imgByte, "唯美少女.jpeg")
	t.Log("[ERROR]: ", err)
}
