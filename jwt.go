package main

import (
	"fmt"
	"os"
	"gopkg.in/square/go-jose.v2/jwt"
)

func main() {
	raw := os.Args[1]
	sharedKey := []byte(os.Getenv("IV_SECRET"))
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		panic(err)
	}

	out := jwt.Claims{}
	if err := tok.Claims(sharedKey, &out); err != nil {
		panic(err)
	}
	err = out.Validate(jwt.Expected{
		Subject: "IV",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
