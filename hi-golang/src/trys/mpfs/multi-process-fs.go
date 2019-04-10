package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	KeyCmdDir = "dir"
)

// @TODO 待实现，用于检测过进行文件检测
// 文件目录的深度
// 扫描树，深度

/**
 * @DATE        2019/4/10
 * @NAME        Joshua Conero
 * @DESCRIPIT   多进程文件系统查看/复制/删除/创建
 **/

// 多进程文件阅读器
type Mpfs struct {
	FileCount int
	DirCount  int
	ProcessX  int
	Dir       string
}

// 树字典
type TreeDickData struct {
	TreePath string // 路径标识符
}
type HashCode string
type TreeDick map[HashCode]TreeDickData

func NewMpfs() *Mpfs {
	mpfs := &Mpfs{}
	if vdir, has := Data[KeyCmdDir]; has {
		mpfs.Dir = vdir
	} else {
	}
	return mpfs
}

func TD2(dir string) TreeDick {
	var td TreeDick
	td = map[HashCode]TreeDickData{}

	// 标准化地址
	dir = strings.Replace(dir, "\\", "/", -1)
	dirLen := len(dir)
	if dirLen > 3 && dir[dirLen-1:] != "/" {
		dir += "/"
	}

	var todoReader func(d1 string)
	todoReader = func(d1 string) {
		d1Src, err := ioutil.ReadDir(d1)
		//fmt.Println("d1 => ", d1, ", err =>", err)
		if err != nil {
			return
		}
		for _, ds := range d1Src {
			pRaw := d1 + ds.Name()
			pCode := strings.Replace(pRaw, dir, "", -1)
			td[HashCode(getHash1(pCode))] = TreeDickData{pCode}

			//fmt.Println(pRaw, pCode, ds.IsDir(), dir)
			// 递归
			if ds.IsDir() {
				pRaw += "/"
				todoReader(pRaw)
			}

		}
	}

	todoReader(dir)
	return td
}

// 命令路由器
func (m *Mpfs) Router() {
	switch action {
	case "info":
		m.Info()
	}
}

// 文件信息
func (m *Mpfs) Info() {
	fmt.Println(TD2(m.Dir))
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

// 获取 hash 码
func getHash1(s string) string {
	h1 := sha1.New()
	return fmt.Sprintf("%x", h1.Sum([]byte(s)))
}

func main() {
	// 测试
	action = "info"
	Data["dir"] = "D:/tmp/"

	mpsf := NewMpfs()
	mpsf.Router()

	fmt.Println(" action ", action)
	fmt.Println(Data)

	//fmt.Println(
	//	getHash1("1"),
	//	" ", getHash1("2"))
}
