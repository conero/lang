package main

// 2018年8月25日 星期六
// 简单服务器
// html{服务器}

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var testSv *http.Server
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

type tSvHd struct {
}

func (sv *tSvHd) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintln(w, r.RequestURI)
	//fmt.Fprintln(res, req)
	name := r.RequestURI
	if name == "" || name == "/"{
		name = "/index.html"
	}
	name = "." + name

	if HasFile(name){
		//http.ServeFile(w, r, r.RequestURI)
		mime := getMime(name)
		if mime != ""{
			w.Header().Set("Content-Type", mime)
		}
		content, err := ioutil.ReadFile(name)
		if err == nil {
			fmt.Fprintln(w, string(content))
		}
	}else {
		fmt.Println("    :( "+ name + "文件不存在！")
	}
}
func getMime(path string) string  {
	idx := strings.LastIndex(path, ".")
	mime := ""
	if idx > -1{
		suffix := strings.ToLower(path[idx+1:])
		switch suffix {
		case "wasm":
			mime = "application/wasm"
		}
	}

	return mime
}

func HasFile(fname string) bool  {
	fh, err := os.Stat(fname)
	if err != nil{
		return false
	}
	return !fh.IsDir()
}
// 服务器初始化
func InitServer()  {
	port, has := data["port"]
	if !has{
		log.Fatal("port=端口号")
	}
	testSv = &http.Server{
		Addr: ":" + port,
		Handler: new(tSvHd),
	}
	fmt.Println("    http://localhost:"+port+"/")
	fmt.Println("    服务器正在运行(测试服务器)")
}

func main() {
	parseData()
	// 监听
	InitServer()
	log.Fatal(testSv.ListenAndServe())
}
