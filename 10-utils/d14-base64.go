package main

import (
	// Go 提供了对 base64 编解码的内建支持
    b64 "encoding/base64"
    "fmt"
)

func main() {

    data := "abc123!?$*&()'-=@~"

	// Go 同时支持标准 base64 以及 URL 兼容 base64。
	// 这是使用标准编码器进行编码，传入一个 []byte
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

	// 解码
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

	// 使用 URL base64 格式进行编解码
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}