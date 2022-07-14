package merge

import (
	"github.com/guolei19850528/go_x_pkgs/image/coder"
	xdraw "golang.org/x/image/draw"
	"image"
	"os"
)

// SimpleMerge
//
// param: name Full path of merged image name 合并后图像名称 全路径
//
// param: imgs image.Image Slice Or Array image.Image 切片或数组
//
// param: r image.Rectangle Uniform zoom Rectangle
//
// param: p mage.Point Uniform zoom Point
func SimpleMerge(name string, imgs []image.Image, r image.Rectangle, p image.Point) error {
	pngCoder := new(coder.PngCoder)
	newImg := image.NewRGBA(r)
	for _, img := range imgs {
		imgScale := image.NewRGBA(r)
		xdraw.CatmullRom.Scale(imgScale, r, img, img.Bounds(), xdraw.Over, nil)
		xdraw.Draw(newImg, r, imgScale, p, xdraw.Over)
	}
	if w, err := os.Create(name); err == nil {
		defer w.Close()
		return pngCoder.Encode(w, newImg, nil)
	}
	return nil
}
