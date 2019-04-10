package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	KeyCmdDir = "dir"
)

// @TODO 待实现，用于检测过进行文件检测
// 文件目录的深度

/**
 * @DATE        2019/4/10
 * @NAME        Joshua Conero
 * @DESCRIPIT   多进程文件系统查看/复制/删除/创建
 **/

type Mpfs struct {
	FileCount int
	DirCount int
	ProcessX int
	Dir string
}

func NewMpfs() *Mpfs  {
	mpfs := &Mpfs{}
	if vdir, has := Data[KeyCmdDir]; has{
		mpfs.Dir = vdir
	}else {
	}
	return mpfs
}
func (m *Mpfs) Router()  {

}

//
func (m *Mpfs) Info()  {

}


var maxProceeX int = 100
var action string
var Data map[string]string

func init() {
	// 命令解析
	osArgsLen := len(os.Args)
	Data = map[string]string{}
	if osArgsLen > 1 {
		action = os.Args[1]
	}
	if osArgsLen > 2 {
		for i, v := range os.Args {
			if i < 2 {
				continue
			}
			idx := strings.Index(v, "=")
			if idx > -1 {
				Data[v[0:idx]] = v[idx+1:]
			}
		}
	}
}


func main() {
	// 测试
	action = "info"
	Data["dir"] = "D:/tmp/"

	mpsf := NewMpfs()
	mpsf.Router()

	fmt.Println(" action ", action)
	fmt.Println(Data)
}
