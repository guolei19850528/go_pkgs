package encrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/guolei19850528/go_pkgs/convert"
)

// Md5 md5 length 32
//
// param: x
// return:
func Md5(x any) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(convert.AnyToString(x))))
}

// Md516 md5 length 16
//
// param: x
// return:
func Md516(x any) string {
	md5Str := Md5(x)
	return md5Str[8:24]
}

// Sha1
//
// param: x
// return:
func Sha1(x any) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(convert.AnyToString(x))))
}

// Sha2256 sha2 256
//
// param: x
// return:
func Sha2256(x any) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(convert.AnyToString(x))))
}
