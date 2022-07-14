package merge_test

import (
	"os"
	"path"
	"testing"
)

var (
	CwdPath, _      = os.Getwd()
	ImageInputPath  = path.Join(CwdPath, "images", "Air", "input")
	ImageOutputPath = path.Join(CwdPath, "images", "Air", "output", "images")
	ImageExtensions = map[string]string{
		".png": ".png",
	}
	PropertyFileExtensions = map[string]string{
		".xlsx": ".xlsx",
	}
)

func TestSimpleMerge(t *testing.T) {
	//imgPath1 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_x_pkgs/image/merge/images/Air/input/a/1.png"
	//imgPath2 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_x_pkgs/image/merge/images/Air/input/c1/1.png"
	//imgPath3 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_x_pkgs/image/merge/images/Air/input/c2/1.png"
	//imgPath4 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_x_pkgs/image/merge/images/Air/input/c3-f/1.png"
	//
	//img1, _ := coder.GetImgByName(imgPath1)
	//img2, _ := coder.GetImgByName(imgPath2)
	//img3, _ := coder.GetImgByName(imgPath3)
	//img4, _ := coder.GetImgByName(imgPath4)
	//
	//imgs := make([]image.Image, 0, 0)
	//imgs = append(imgs, img1)
	//imgs = append(imgs, img2)
	//imgs = append(imgs, img3)
	//imgs = append(imgs, img4)
	//slice.ReverseNoSort(imgs)
	//merge.SimpleMerge(path.Join(ImageOutputPath, "1.png"), imgs, image.Rect(0, 0, 2200, 2200), image.Pt(0, 0))
}
