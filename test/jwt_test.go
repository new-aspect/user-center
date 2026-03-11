package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(*testing.T) {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	fmt.Printf("foo: %v\n", claims.Foo)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(mySigningKey)
	fmt.Println(signedString, err)
}

func TestParseJWT(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE3NzMzMDUyNjUsIm5iZiI6MTc3MzIxODg2NSwiaWF0IjoxNzczMjE4ODY1fQ.Zx3JJrSyz34ryxr8OrB3wDIIk4erVR3_rCYYoup18QI"

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		panic(err.Error())
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		fmt.Println(claims.Foo, claims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
}
