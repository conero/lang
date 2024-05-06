package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

/**
 * @DATE        2019/5/17
 * @NAME        Joshua Conero
 * @DESCRIPIT   描述 descript
 **/

// QQ 邮箱发送示例
func sendQQCom() {
	// conero@qq.com
	// POP3/SMTP服务
	pswd := "授权码"
	auth := smtp.PlainAuth("", "user@qq.com", pswd, "smtp.qq.com")
	to := []string{"toname@163.com"}
	nickname := "Jc"
	user := "xxxxx@qq.com"
	subject := "test mail"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " +
		subject + "\r\n" + content_type + "\r\n\r\n" + body)

	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	} else {
		fmt.Println("已发送邮箱！")
	}
}

// xml 数据结构
type xDef struct {
	Nickname string `xml:"nickname,attr"`
}
type xEmail struct {
	Id       string `xml:"id,attr"`
	Account  string `xml:"account,attr"`
	Host     string `xml:"host,attr"`
	Authcode string `xml:"authcode,attr"`
	Def      xDef   `xml:"default"`
}

type X struct {
	XMLName xml.Name `xml:"SmptAuth"`
	Email   xEmail   `xml:"email"`
}

func getX() X {
	fname := "./smtp_auth.xml"
	ct, err := ioutil.ReadFile(fname)

	// 测试脚本
	if err != nil {
		fname = "D:/conero/gits/conero/lang/hi-lang/hi-golang/src/mail/base1/smtp_auth.xml"
		ct, err = ioutil.ReadFile(fname)
	}

	if err != nil {
		fmt.Println("配置文件不存在：" + err.Error())
	}

	var x X
	if err := xml.Unmarshal(ct, &x); err != nil {
		log.Fatal(err)
	}
	return x
}

// 获取模板字符串
func getFormatXml() string {
	var x X = X{
		Email: xEmail{
			Id:       "类型",
			Account:  "setting@mail.com",
			Host:     "smtp.mail.com",
			Authcode: "IamThEGuYWIthBigDREAMInThEChINA.",
			Def: xDef{
				Nickname: "CCW#.3co",
			},
		},
	}
	ss, err := xml.Marshal(x)
	if err != nil {
		return "E111/" + err.Error()
	}
	return `<?xml version="1.0" encoding="UTF-8"?>` + string(ss)
}

var gX X

const cmdOutputXml = "--xml"

var cmdOutputXmlMk bool = false

func SendEmail(to, subject, content, user, nickname string) error {
	email := gX.Email
	if nickname != "" {
		nickname = email.Def.Nickname
	}
	auth := smtp.PlainAuth("", email.Account, email.Authcode, email.Host)
	content_type := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " +
		subject + "\r\n" + content_type + "\r\n\r\n" + content)
	return smtp.SendMail("smtp.qq.com:25", auth, user, strings.Split(to, ","), msg)
}

var Data map[string]string

var checkKeyValid bool = false // key 有效性检查

// 控制台
func main() {
	if cmdOutputXmlMk {
		erro := ioutil.WriteFile("./smtp_auth.create-exaple.xml", []byte(getFormatXml()), 0777)
		if erro != nil {
			fmt.Println(" 错误： " + erro.Error())
		}
		return
	}

	// 右键发送
	if !checkKeyValid {
		fmt.Println("$ * t=接收人/to s=主题/subject c=内容/content n=名称/name")
		fmt.Println("  * 参数无效，缺失或者为空！")
		fmt.Printf("	. 输入参数：%v \r\n", Data)
		fmt.Println()
		fmt.Println("  * 说明")
		fmt.Println("	. to 支持多人如 'to=a1,a2'")
		fmt.Println("	. " + cmdOutputXml + " 生成xml配置模板")
		return
	}
	err := SendEmail(Data["t"], Data["s"], Data["c"], gX.Email.Account, Data["n"])
	if err != nil {
		fmt.Println(" ） 邮箱发送失败！")
		fmt.Println(" ） " + err.Error())
	} else {
		fmt.Println("E000/邮箱已发送")
	}
}

func init() {
	gX = getX()
	Data = map[string]string{}

	for i := 1; i < len(os.Args); i++ {
		ss := os.Args[i]
		if ss == cmdOutputXml {
			cmdOutputXmlMk = true
		}
		idx := strings.Index(ss, "=")
		if idx > 0 {
			key := ss[0:idx]
			value := ss[idx+1:]
			Data[key] = value
		}
	}

	keys := []string{"t", "s", "c", "n"}
	tmpValid := true
	for _, k := range keys {
		if _, has := Data[k]; !has {
			tmpValid = false
			break
		}
	}
	checkKeyValid = tmpValid
}
