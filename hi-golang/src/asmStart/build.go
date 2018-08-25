package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

// 简单 wasm 编译器


var data map[string]string

func parseData(){
	data = map[string]string{}
	for _,v := range os.Args{
		v = strings.TrimSpace(v)
		idx := strings.Index(v, "=")
		if idx > -1{
			key := v[0:idx]
			data[key] = v[idx+1:]
		}
	}
	//fmt.Println(os.Args[0], path.Dir(os.Args[0]))
}

func inArray(s string, a []string) int  {
	idx := -1
	for i, v := range a{
		if s == v{
			idx = i
			break
		}
	}
	return  idx
}

func buildWasm()  {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	ignore, has := data["ignore"]
	ignoreQue := []string{}
	if has{
		ignoreQue = strings.Split(ignore, ",")
	}
	for _, file := range files {
		name := file.Name()
		if strings.ToLower(path.Ext(name)) != ".go"{
			continue
		}
		if has && inArray(name, ignoreQue) > -1{
			continue
		}
		name1 := strings.Replace(name, path.Ext(name), "", -1)
		cmdStr := " go build -o "+name1+".wasm "+name1+".go"
		fmt.Println("   $>" + cmdStr)
		//cmd := exec.Command("go", "build", "-0", name1+".wasm", name1+".go")
		cmd := exec.Command("go", "build -0 "+name1+".wasm "+ name1+".go")
		//cmd := exec.Command("ping", "conero.cn")
		if err := cmd.Run(); err != nil {
			fmt.Println("    错误： "+ err.Error())
		}
	}
}

func main() {
	parseData()
	buildWasm()
}
