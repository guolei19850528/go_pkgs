package dirtree

import (
	"io/fs"
	"os"
	"path"
	"strings"
)

// TreeNode
//Directory Tree Node
type TreeNode struct {
	//node name example:"1.png"
	Name string

	//node base name does not contain extension example:"1"
	BaseName string

	//node extension example:".png"
	ExtName string

	//current dir name example: "a"
	CurrDirName string

	//base dir path example:"/Users/guolei/images"
	BaseDirPath string

	//is dir
	IsDir bool

	//full path example:"/Users/guolei/images/a/1.png"
	FullPath string

	//file info
	FileInfo os.FileInfo

	//file mode
	Type fs.FileMode

	//array of subordinate nodes
	Children []*TreeNode
}

// Tree
//
// param: name dir path
// param: matchExt match extension map
// param: sortFunc sort function
// return:
func Tree(name string, matchExt map[string]string, sortFunc func([]*TreeNode)) []*TreeNode {
	treeNodes := make([]*TreeNode, 0, 0)
	dirEntries, _ := os.ReadDir(name)
	for _, dirEntry := range dirEntries {
		treeNode := new(TreeNode)
		treeNode.Name = dirEntry.Name()
		treeNode.ExtName = path.Ext(treeNode.Name)
		treeNode.BaseName = strings.TrimSuffix(treeNode.Name, treeNode.ExtName)
		treeNode.CurrDirName = path.Base(name)
		treeNode.BaseDirPath = path.Dir(name)
		treeNode.FullPath = path.Join(name, treeNode.Name)
		treeNode.IsDir = dirEntry.IsDir()
		treeNode.Type = dirEntry.Type()
		treeNode.FileInfo, _ = dirEntry.Info()
		if treeNode.IsDir {
			treeNode.Children = Tree(treeNode.FullPath, matchExt, sortFunc)
		} else {
			if matchExt != nil {
				if _, ok := matchExt[treeNode.ExtName]; !ok {
					continue
				}
			}
		}
		treeNodes = append(treeNodes, treeNode)
	}
	if sortFunc != nil {
		sortFunc(treeNodes)
	}
	return treeNodes
}
