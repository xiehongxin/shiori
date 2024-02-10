package pkg

import (
	"fmt"
	"testing"
)

func TestGoogleAuth(t *testing.T) {
	fmt.Println("----------------- generate secret -------------------")
	secret := GetSecret()
	fmt.Println("secret:" + secret)

	fmt.Println("----------------- google code verify ----------------------")
	var code int32
	fmt.Print("Please input Google Auth Code:")
	for {
		_, err := fmt.Scan(&code)
		if err == nil {
			break
		}

		fmt.Print("Input error, please re-enter:")
	}

	b := VerifyCode(secret, code)
	if b {
		fmt.Println("verify success！")
	} else {
		fmt.Println("verify failed！")
	}
}
