package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		if os.Args[1] == "encode" {
			encodedString := base64.StdEncoding.EncodeToString([]byte(os.Args[2]))
			fmt.Println("encodedString:", encodedString)
		} else if os.Args[1] == "decode" {
			decodeString, err := base64.StdEncoding.DecodeString(os.Args[2])
			if err != nil {
				log.Println("decode string err:", err)
				return
			}
			fmt.Printf("decodeString: %s", decodeString)
		} else {
			log.Println("first arg must be encode or decode")
		}
	} else {
		log.Println("args must equal 3 ")
	}
}
