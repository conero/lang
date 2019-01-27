package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/str"
)

func EmptyAction() {
	fmt.Println(" go 加密库学测试：")
	fmt.Println(" $ cmd [options]")
	fmt.Println(" $ en   加密")
	fmt.Println(" $ de   解密")
}

type BaseCmds struct {
}

// 加密
func (c BaseCmds) EnApp() {
}

// 解密
func (c BaseCmds) DeApp() {
}

func (c BaseCmds) BadEnter(s string) {
	fmt.Println(" 命名无法识别[" + s + "]！")
}

var baseCmds BaseCmds
var app bin.App

const (
	CmdMd5En = "md5"
	CmdShaEn = "sha1"
	CmdAllBase = "all"
	CmdEn = "en"
	CmdDe = "de"

)

func RouterHundler(opt string)  {
	app = bin.GetApp()
	switch {
	case str.InQuei(opt, []string{CmdEn}) > -1:
		baseCmds.EnApp()
	case str.InQuei(opt, []string{CmdMd5En}) > -1:			// md5
		txt := app.Next(CmdMd5En)
		fmt.Printf(" (md5)明文： %s\n", txt)
		fmt.Printf(" %s ", MdsEnApp(txt))
	case str.InQuei(opt, []string{CmdShaEn}) > -1:			// sha1
		txt := app.Next(CmdShaEn)
		fmt.Printf(" (sha1)明文： %s\n", txt)
		fmt.Printf(" %s ", Sha1EnApp(txt))
	case str.InQuei(opt, []string{CmdAllBase}) > -1:			// all
		txt := app.Next(CmdAllBase)
		fmt.Printf(" 明文： %s\n", txt)
		fmt.Printf(" md5: %s\n", MdsEnApp(txt))
		fmt.Printf(" sha1: %s\n", Sha1EnApp(txt))
		fmt.Printf(" sha256: %s\n", sha256EnApp(txt))
		fmt.Printf(" sha512: %s\n", sha512EnApp(txt))
	case str.InQuei(opt, []string{CmdShaEn}) > -1:
		baseCmds.DeApp()
	default:
		baseCmds.BadEnter(opt)
	}
}

func main() {
	router := &bin.Router{
		EmptyAction:    EmptyAction,
		UnfindAction: RouterHundler,
	}
	bin.Adapter(router)
	bin.Run()
	//fmt.Println(aesEnTxt("gfgfgfgdsddksdsdjkj", sha256EnApp("yh")))
	//
	//// GCM
	//ExampleNewGCM_encrypt()
	//ExampleNewGCM_decrypt()
	//
	//// CBC
	//ExampleNewCBCEncrypter()
	//ExampleNewCBCDecrypter()
	//
	//// CFB
	//ExampleNewCFBEncrypter()
	//ExampleNewCFBDecrypter()
	//
	//// CTR
	//ExampleNewCTR()
	//// OFB
	//ExampleNewOFB()
	//
	//ExampleStreamWriter()
	//ExampleStreamReader()
}
