package main

import (
	"log"

	"github.com/x7c7v7i87/tools/authcode"
)

func main() {
	log.Println("test....")
	key := "abcddddddd"
	str, _ := authcode.Encrypt("ewewewefwefwefew", key, 0)
	log.Println(str)
	log.Println(authcode.Decrypt(str, key))
}
