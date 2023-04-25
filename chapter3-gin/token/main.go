package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	// 创建秘钥
	key := []byte("aaa")
	// 创建Token结构体
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "zhangshan",
		"pass": "123123",
	})
	// 调用加密方法，发挥Token字符串
	signingString, err := claims.SignedString(key)
	if err != nil {
		return
	}
	fmt.Println(signingString)

	// 根据Token字符串解析成Claims结构体
	var data = jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(signingString, data, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("Header: %v\n", token.Header)
		fmt.Printf("Signature: %v\n", token.Signature)
		fmt.Printf("Raw: %v\n", token.Raw)
		fmt.Printf("err: %v\n", token.Claims.Valid())
		fmt.Printf("err: %v\n", token.Claims.Valid() == nil)
		fmt.Printf("输出: %v\n", token.Claims)

		// var res jwt.MapClaims
		// if claims, ok := token.Claims.(res); ok && token.Valid {
		// 	return claims, nil
		// }

		return []byte("aaa"), nil
	})
	fmt.Printf("输出d: %v\n", data["pass"])
	fmt.Printf("输出d: %v\n", data["user"])

	if err != nil {
		fmt.Println(err)
		return
	}

}

//这边是输出结果
// &{ 0xc0000c2690 map[alg:ES256 typ:JWT] map[user:zhangshan]  false}
// 这是加密后的字符串
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoiMTIzMTIzIiwidXNlciI6InpoYW5nc2hhbiJ9.-2-xIJXMGKV-GyhM24OKbDVqWs4dsIANBsGhzXEfEFM
