package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	key := Key()
	fmt.Println(key)
	fmt.Println(newIdKey())
}

func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", buf)
}

func newIdKey() int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	return n.Int64()
}
