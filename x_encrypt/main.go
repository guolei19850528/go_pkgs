package x_encrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

// Md5 md5 length 32
//
// param: x
// return:
func Md5(x any) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(x.(string))))
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
	return fmt.Sprintf("%x", sha1.Sum([]byte([]byte(x.(string)))))
}

// Sha2256 sha2 256
//
// param: x
// return:
func Sha2256(x any) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte([]byte(x.(string)))))
}
