package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// 迷宫
type Maze struct {
	TargetDir string // 目标地址
	innerTask func() // 内部任务
	MnChan    chan MazeNode
}

// 迷宫节点
type MazeNode struct {
	FileCount int
	DirCount  int
	PathQueue []string
	Path      string
	IsDir     bool
}

/**
任务设置
*/
func (m *Maze) Task(todoTask func()) *Maze {
	m.innerTask = todoTask
	return m
}

/**
默认任务
*/
func (m *Maze) defTask() {
}

/**
调度中心
*/
func (m *Maze) dispatch(vpath string) {
	tracks, err := ioutil.ReadDir(vpath)
	// 错误
	if err == nil {
		dirQue := []string{}
		fileCtt := 0 // 文件统计

		mNode := MazeNode{
			IsDir: false,
			Path:  vpath,
		}

		for _, track := range tracks {
			name := track.Name()
			vpath := m.TargetDir + name

			if track.IsDir() {
				dirQue = append(dirQue, vpath)
				mNode.IsDir = true
			} else {
				fileCtt += 1
			}
		}
		fmt.Println(dirQue)
		//fmt.Println(mNode)

		//m.MnChan <- mNode

		// 线程分发
		for _, p := range dirQue {
			go m.dispatch(p)
		}

		m.MnChan <- mNode
	}
}

/**
寻物，迷宫入口。并获取字符信息
*/
func (m *Maze) Search() string {
	var msg string
	if m.MnChan == nil{
		m.MnChan = make(chan MazeNode)
	}
	_, err := ioutil.ReadDir(m.TargetDir)
	if err != nil {
		msg = err.Error()
	} else {
		m.dispatch(m.TargetDir)
	}
	return msg
}

//集合
func (m *Maze) Gather() {
	// 线程保持
	for mn := range m.MnChan {
		fmt.Println(mn.Path)
	}
}

/**
迷宫示例
*/
func NewMaze(tgtDir string) *Maze {
	return &Maze{TargetDir: getStdDir(tgtDir), innerTask: nil}
}

/**
标准路径字符串
*/
func getStdDir(dirName string) string {
	dirName = strings.TrimSpace(dirName)
	dirName = strings.ReplaceAll(dirName, "\\", "/")
	reg := regexp.MustCompile("[/]{2,}")
	dirName = reg.ReplaceAllString(dirName, "/")

	vlen := len(dirName)
	if dirName[vlen-1:] != "/" {
		dirName += "/"
	}
	return dirName
}
