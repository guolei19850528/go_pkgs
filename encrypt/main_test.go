package encrypt_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/encrypt"
	"testing"
)

func Test(t *testing.T) {
	var encryptStr string
	encryptStr = encrypt.Md5(123456)
	fmt.Println("Md5:", encryptStr)
	encryptStr = encrypt.Md516(123456)
	fmt.Println("Md516:", encryptStr)
	encryptStr = encrypt.Sha1(123456)
	fmt.Println("Sha1:", encryptStr)
	encryptStr = encrypt.Sha2256(123456)
	fmt.Println("Sha256:", encryptStr)

}
