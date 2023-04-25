package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TokenExpireDuration = 2 * time.Hour //过期时间
var hmacSampleSecret = []byte("aaa")

type AuthClaim struct {
	UID int64 `json:"uid"`
	jwt.StandardClaims
}

func New(uid int64) (tokenStr string) {
	var authClaim AuthClaim
	authClaim.UID = uid
	authClaim.StandardClaims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaim)
	tokenString, _ := token.SignedString(hmacSampleSecret) //私钥加密
	return tokenString
}

func Parse(tokenString string) (auth AuthClaim, Valid bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	Valid = token.Valid //token是否有效 true有效  false无效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && Valid {
		fmt.Println(claims)
		auth.UID = int64(claims["uid"].(float64))       //自定义的UID
		auth.ExpiresAt = int64(claims["exp"].(float64)) //过期时间
	}
	return
}
func main() {
	str := New(1)
	fmt.Println("生成", str)

	auth, err := Parse(str)
	if !err {
		fmt.Println("出错了", err)
		return
	}
	fmt.Printf("解析%+v", auth)
}

//这边是输出结果
// &{ 0xc0000c2690 map[alg:ES256 typ:JWT] map[user:zhangshan]  false}
// 这是加密后的字符串
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoiMTIzMTIzIiwidXNlciI6InpoYW5nc2hhbiJ9.-2-xIJXMGKV-GyhM24OKbDVqWs4dsIANBsGhzXEfEFM
