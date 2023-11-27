package test

import (
	"fintechpractices/tools"
	"fmt"
	"testing"
	"time"
)

func TestMd5(t *testing.T) {
	m1 := tools.GenMD5WithSalt("hello world", tools.Salt)
	m2 := tools.GenMD5("hello worldCC-fintech-practices")
	if m1 == m2 {
		fmt.Println(m1, len(m1))
	} else {
		t.Errorf("%s not equal to %s", m1, m2)
	}
}

func TestEncrypt(t *testing.T) {
	tools.TestEncrypt()
}

func TestParseToken(t *testing.T) {
	token, err := tools.GenToken("admin")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(token)

	time.Sleep(time.Second)
	claims, err := tools.ParseToken(token)
	if err != nil {
		t.Error(err.Error())
	}
	if claims.UserAccount != "admin" {
		t.Fail()
	} else {
		fmt.Println(claims)
	}
}
