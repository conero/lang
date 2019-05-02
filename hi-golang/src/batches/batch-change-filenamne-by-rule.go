/**
2019年5月1日 星期三
批量文件名称修改
*/

package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const BatchChangeFileBaseDirName = "jc___batch-maker-files"

func todoRun(dirx, targetDir string) {
	cttMap := map[string]int{
		"dir": 0,
		"img": 0,
	}
	if dirx == "" {
		dirx = "./"
	}
	if targetDir == "" {
		targetDir = "./" + BatchChangeFileBaseDirName + "/"
	}

	if fs.IsDir(targetDir) {
		os.RemoveAll(targetDir)
	} else {
		fs.CheckDir(targetDir)
	}
	fis, err := ioutil.ReadDir(dirx)
	if err != nil {
		log.Fatal(err.Error())
	}
	compStr := map[string]string{
		"_b.": "back",
		"_d.": "bottom",
		"_f.": "front",
		"_l.": "left",
		"_r.": "right",
		"_u.": "top",
	}
	for _, fi := range fis {
		dirPath := dirx + fi.Name()
		if fi.IsDir() {
			cttMap["dir"] += 1
			cfis, _ := ioutil.ReadDir(dirPath)
			for _, cfi := range cfis {
				if cfi.IsDir() {
					continue
				}
				imgName := cfi.Name()
				oldfile := dirPath + "/" + imgName
				//fmt.Println(imgUrl)

				for vs, ts := range compStr {
					idx := strings.Index(imgName, vs)
					if idx > -1 {
						name := fi.Name()
						newPath := targetDir + name + "/"
						fs.CheckDir(newPath)
						newfile := newPath + name + "." + ts + "." + imgName[idx+len(vs):]
						//fmt.Println(newfile, oldfile)
						fs.Copy(newfile, oldfile)
						cttMap["img"] += 1
					}
				}
			}
		}
	}
	fmt.Println("源地址：" + dirx)
	fmt.Println("目标地址：" + targetDir)
	fmt.Printf(" 遍历目录 %v，复制文件： %v\r\n", cttMap["dir"], cttMap["img"])

	// 删除空目录
	if cttMap["img"] == 0 {
		if fs.IsDir(targetDir) {
			os.RemoveAll(targetDir)
		}
	}
}

func main() {
	bin.Adapter(&bin.Router{
		EmptyAction: func() {
			doc := "****** 配置 *********\r\n" +
				"  run      						    \r\n" +
				"    src=./								源地址 \r\n" +
				"    tgt=./jc___batch-maker-files/		目标地址 \r\n"
			fmt.Println(doc)
		},
		FuncAction: func(action string, a *bin.App) bool {
			if action == "run" {
				//fmt.Println(a.Data)
				var vv interface{}
				var has bool
				var src, tgt string
				vv, has = a.Data["src"]
				if has {
					src = vv.(string)
				}

				vv, has = a.Data["tgt"]
				if has {
					tgt = vv.(string)
				}

				todoRun(src, tgt)
				return true
			}
			return false
		},
	})
	bin.Run()
}
