package main

import (
	"fmt"
	"github.com/conero/uymas/bin"
	"github.com/conero/uymas/str"
)

const (
	CmdMd5En   = "md5"
	CmdShaEn   = "sha1"
	CmdAllBase = "all"
	CmdAESCBC  = "aes"
	CmdRsaOaep = "rsa"
	CmdEn      = "en"
	CmdDe      = "de"
	DefKey     = "Joshua Conero"
)

func EmptyAction() {
	fmt.Println(" go 加密库学测试：")
	fmt.Println(" $ cmd [options]")
	fmt.Println(" $ md5 <text> 文本编码，其他[sha1, all]类似")
	fmt.Println(" $ aes <text> [options] aes 加密测试(rsa类似)")
	fmt.Println(" 		--k=keys   指定加密秘钥")
	fmt.Println(" 		--de       跟换为解密")
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

func RouterHundler(opt string) {
	app = bin.GetApp()
	switch {
	case str.InQuei(opt, []string{CmdEn}) > -1:
		baseCmds.EnApp()
	case str.InQuei(opt, []string{CmdMd5En}) > -1: // md5
		txt := app.Next(CmdMd5En)
		fmt.Printf(" (md5)明文： %s\n", txt)
		fmt.Printf(" %s ", MdsEnApp(txt))
	case str.InQuei(opt, []string{CmdShaEn}) > -1: // sha1
		txt := app.Next(CmdShaEn)
		fmt.Printf(" (sha1)明文： %s\n", txt)
		fmt.Printf(" %s ", Sha1EnApp(txt))
	case str.InQuei(opt, []string{CmdAESCBC}) > -1:
		k, isK := app.Data["k"]
		key := DefKey
		if isK {
			if _, isString := k.(string); isString {
				key = k.(string)
			}
		}
		txt := app.Next(CmdAESCBC)

		var ctxt string
		txtType := "明文"
		if app.HasSetting("de") {
			ctxt = rsaDe(txt, key)
			txtType = "密文"
			txt = "..."
		} else {
			ctxt = aesEnTxt(txt, key)
		}

		fmt.Printf(" (aes)%s： %s\n", txtType, txt)
		fmt.Printf(" (aes)秘钥： %s\n", key)
		fmt.Printf(" %s ", ctxt)
	case str.InQuei(opt, []string{CmdRsaOaep}) > -1:
		k, isK := app.Data["k"]
		key := string(rsaLabel)
		if isK {
			if _, isString := k.(string); isString {
				key = k.(string)
			}
		}
		txt := app.Next(CmdRsaOaep)

		var ctxt string
		txtType := "明文"
		if app.HasSetting("de") {
			ctxt = rsaDe(txt, key)
			txtType = "密文"
			txt = "..."
		} else {
			ctxt = rsaEn(txt, key)
		}

		fmt.Printf(" (rsa)%s： %s\n", txtType, txt)
		fmt.Printf(" (rsa)秘钥： %s\n", key)
		fmt.Printf(" %s ", ctxt)

	case str.InQuei(opt, []string{CmdAllBase}) > -1: // all
		txt := app.Next(CmdAllBase)
		fmt.Printf(" 明文： %s\n", txt)
		fmt.Printf(" md5: %s\n", MdsEnApp(txt))
		fmt.Printf(" sha1: %s\n", Sha1EnApp(txt))
		fmt.Printf(" sha256: %s\n", sha256EnApp(txt))
		fmt.Printf(" sha512: %s\n", sha512EnApp(txt))
		fmt.Printf("\n")
		fmt.Printf(" aes(key=%s): %s\n", DefKey, aesEnTxt(txt, DefKey))
	case str.InQuei(opt, []string{CmdShaEn}) > -1:
		baseCmds.DeApp()
	default:
		baseCmds.BadEnter(opt)
	}
}

func main() {
	router := &bin.Router{
		EmptyAction:  EmptyAction,
		UnfindAction: RouterHundler,
	}
	bin.Adapter(router)
	bin.Run()
	//fmt.Println(aesEnTxt("dududjdjhdsh 杨华", "yh"))
	//fmt.Println(aesDnTxt(aesEnTxt("dududjdjhdsh 杨华", "yh"), "yh"))

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
