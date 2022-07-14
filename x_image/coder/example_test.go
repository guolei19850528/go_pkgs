package coder_test

import (
	"fmt"
	"github.com/guolei19850528/go_x_pkgs/image/coder"
)

func ExampleNewCoder() {
	name := "1.png"
	imageCoder, err := coder.NewCoder(name)
	if err == nil {
		fmt.Println(imageCoder)
	}
}

func ExampleGetImgByName() {
	name := "1.png"
	img, err := coder.GetImgByName(name)
	if err == nil {
		fmt.Println(img)
	}
}
