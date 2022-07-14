package x_coder_test

import (
	"fmt"
	"github.com/guolei19850528/go_x_pkgs/image/x_coder"
)

func ExampleNewCoder() {
	name := "1.png"
	imageCoder, err := x_coder.NewCoder(name)
	if err == nil {
		fmt.Println(imageCoder)
	}
}

func ExampleGetImgByName() {
	name := "1.png"
	img, err := x_coder.GetImgByName(name)
	if err == nil {
		fmt.Println(img)
	}
}
