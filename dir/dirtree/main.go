package dirtree

import (
	"io/fs"
	"os"
	"path"
	"strings"
)

// TreeNode
//Directory Tree Node
//
//目录树节点
type TreeNode struct {
	//node name 节点名称
	Name string

	//node base name does not contain extension 节点名称 不包含扩展名
	BaseName string

	//node extension 节点扩展名
	ExtName string

	//current dir name 当前节点所在文件夹名称
	CurrDirName string

	//base dir 当前节点所在文件夹路径地址 不包含当前节点所在文件夹
	BaseDirPath string

	//is dir 是否文件夹
	IsDir bool

	//full path 完整路径
	FullPath string

	//file info os.DirEntry.Info()
	FileInfo os.FileInfo

	//file mode os.DirEntry.Type()
	Type fs.FileMode

	//array of subordinate nodes  子节点数组
	Children []*TreeNode
}

// Tree
//
//Get directory tree recursively 递归获取目录树
//
// param: name dir path 目录路径地址
//
// param: matchExt match extension map nil all files 匹配扩展名Map
//
// param: sortFunc sort function nil no sort 排序函数
//
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
