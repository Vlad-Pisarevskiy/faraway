package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {

	challenge := "48a4bc113de"
	nonce := 1

	data := challenge + strconv.Itoa(nonce)
	value := sha256.Sum256([]byte(data))

	fmt.Println(hex.EncodeToString(value[:]))
}
