package encrypt_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/encrypt"
)

func ExampleMd5() {
	str := encrypt.Md5("123456")
	fmt.Println(str)
	//Output:
	//e10adc3949ba59abbe56e057f20f883e
}

func ExampleMd16() {
	str := encrypt.Md516("123456")
	fmt.Println(str)
	//Output:
	//49ba59abbe56e057
}

func ExampleSha1() {
	str := encrypt.Sha1("123456")
	fmt.Println(str)
	//Output:
	//7c4a8d09ca3762af61e59520943dc26494f8941b
}

func ExampleSha2256() {
	str := encrypt.Sha2256("123456")
	fmt.Println(str)
	//Output:
	//8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
}
