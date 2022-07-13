package merge_test

import (
	"github.com/guolei19850528/go_pkgs/img/coder"
	"os"
	"path"
	"testing"
)

var (
	CwdPath, _      = os.Getwd()
	ImageInputPath  = path.Join(CwdPath, "images", "Air", "input")
	ImageOutputPath = path.Join(CwdPath, "images", "Air", "output")
	ImageExtensions = map[string]string{
		".png": ".png",
	}
	PropertyFileExtensions = map[string]string{
		".xlsx": ".xlsx",
	}
)

func TestSimpleMerge(t *testing.T) {
	imgPath1 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_pkgs/img/merge/images/Air/input/a/1.png"
	imgPath2 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_pkgs/img/merge/images/Air/input/c1/1.png"
	imgPath3 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_pkgs/img/merge/images/Air/input/c2/1.png"
	imgPath4 := "/Users/guolei/Dev/JetBrains/GoLand/guolei-projects/go_pkgs/img/merge/images/Air/input/c3-f/1.png"

	coder.GetImgByName(imgPath1)
	coder.GetImgByName(imgPath2)
	coder.GetImgByName(imgPath3)
	coder.GetImgByName(imgPath4)

	//merge.SimpleMerge("a.png")
}
