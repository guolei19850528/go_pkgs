package x_encrypt_test

import (
	"fmt"
	"github.com/guolei19850528/go_x_pkgs/encrypt"
	"testing"
)

func Test(t *testing.T) {
	var encryptStr string
	encryptStr = x_encrypt.Md5(123456)
	fmt.Println("Md5:", encryptStr)
	encryptStr = x_encrypt.Md516(123456)
	fmt.Println("Md516:", encryptStr)
	encryptStr = x_encrypt.Sha1(123456)
	fmt.Println("Sha1:", encryptStr)
	encryptStr = x_encrypt.Sha2256(123456)
	fmt.Println("Sha256:", encryptStr)

}
