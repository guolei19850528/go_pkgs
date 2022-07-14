package x_merge_test

import (
	"fmt"
	"github.com/guolei19850528/go_x_pkgs/image/x_coder"
	"github.com/guolei19850528/go_x_pkgs/image/x_merge"
	"github.com/guolei19850528/go_x_pkgs/x_slice"
	"image"
	"path"
)

func ExampleSimpleMerge() {

	imgPath1 := "image path"
	imgPath2 := "image path"
	imgPath3 := "image path"
	imgPath4 := "image path"

	img1, _ := x_coder.GetImgByName(imgPath1)
	img2, _ := x_coder.GetImgByName(imgPath2)
	img3, _ := x_coder.GetImgByName(imgPath3)
	img4, _ := x_coder.GetImgByName(imgPath4)

	imgs := make([]image.Image, 0, 0)
	imgs = append(imgs, img1)
	imgs = append(imgs, img2)
	imgs = append(imgs, img3)
	imgs = append(imgs, img4)
	x_slice.ReverseNoSort(imgs)
	err := x_merge.SimpleMerge(path.Join("your output path", "1.png"), imgs, image.Rect(0, 0, 2200, 2200), image.Pt(0, 0))
	if err == nil {
		fmt.Println("successful")
	}
}
