package main

import (
	"fmt"
	"go-guide/lib/jwt/jwt1/jt"
)

func main() {
	token, err := jt.CreateToken(1)
	fmt.Println("token:", token, ";err:", err)

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTYzNzE2MTYzMCwiaXNzIjoiYmluZ28ifQ.OTj1LZ5Pdy_Rkgv5w3ruL3n3I68EhW3k_tFk3-CiFPo"
	clamis, err := jt.ParseToken(token)
	fmt.Printf("clamis:%+v\n", clamis)
	fmt.Println("err:", err)
}
