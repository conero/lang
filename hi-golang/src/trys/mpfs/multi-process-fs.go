package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"
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
	Depth    int    // 深度
	IsDir    bool
	Size     int64
	ModTime  time.Time // 修改时间
}
type HashCode string
type TreeDick map[HashCode]TreeDickData

// info 返回的数据格式
type MpInfoData struct {
	Td        TreeDick // 数据字典
	DirCount  int
	FileCount int
	Size      int64
	Dir       string
	Depth     int
	RunTimes  string
}

// 构造函数
func NewMpfs() *Mpfs {
	mpfs := &Mpfs{}
	if vdir, has := Data[KeyCmdDir]; has {
		mpfs.Dir = vdir
	} else {
	}
	return mpfs
}

// 简单树形字典生成器
// 单线程数据统计
func TD2(dir string) MpInfoData {
	rTm := getTime()
	var infoDt MpInfoData = MpInfoData{
		DirCount:  0,
		FileCount: 0,
		Size:      0,
	}
	var td TreeDick
	td = map[HashCode]TreeDickData{}

	// 标准化地址
	dir = strings.Replace(dir, "\\", "/", -1)
	infoDt.Dir = dir
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
			vdepth := strings.Count(pCode, "/") + 1
			td[HashCode(getHash1(pCode))] = TreeDickData{
				TreePath: pCode,
				Depth:    vdepth,
				IsDir:    ds.IsDir(),
				Size:     ds.Size(),
				ModTime:  ds.ModTime(),
			}
			infoDt.Depth = int(math.Max(float64(infoDt.Depth), float64(vdepth)))
			//fmt.Println(pRaw, pCode, ds.IsDir(), dir)
			// 递归
			if ds.IsDir() {
				infoDt.DirCount += 1
				pRaw += "/"
				todoReader(pRaw)
			} else {
				infoDt.FileCount += 1
				infoDt.Size += ds.Size()
			}

		}
	}

	todoReader(dir)
	infoDt.Td = td
	infoDt.RunTimes = rTm()
	return infoDt
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
	td2 := TD2(m.Dir)
	//fmt.Println(td2)
	tClk := getTime()
	js, _ := json.Marshal(td2)
	fmt.Println(string(js))
	fmt.Println(tClk())
}

var maxProceeX int = 100 // 最大的进程数
var action string
var Data map[string]string
var globalTmer func() string


func init() {
	globalTmer = getTime()
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

// 获取时间差
func getTime() func() string {
	tmStart := time.Now()
	return func() string {
		d := time.Since(tmStart)
		return fmt.Sprintf("%s", d)
	}
}

func main() {
	// 测试
	action = "info"
	Data["dir"] = "D:/tmp/"
	//Data["dir"] = "D:/"
	Data["dir"] = "D:/tmp"

	// 入口配置
	mpsf := NewMpfs()
	mpsf.Router()

	fmt.Println(" action ", action)
	fmt.Println(Data)

	//fmt.Println(
	//	getHash1("1"),
	//	" ", getHash1("2"))

	// 时间
	fmt.Println(".  程序运行的全部时间：", globalTmer())
}
