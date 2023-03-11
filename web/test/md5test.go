package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("hello")
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(hex.EncodeToString(cipherStr))

}
