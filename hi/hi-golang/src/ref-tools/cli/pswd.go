package main

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
)

// 入口
func main() {
	showPswd()
}

// 密码框输入
func showPswd() {
	fmt.Printf("Password: ")

	// Silent. For printing *'s use gopass.GetPasswdMasked()
	pass, err := gopass.GetPasswd()
	if err != nil {
		// Handle gopass.ErrInterrupted or getch() read error
		fmt.Println(err.Error())
		return
	}

	fmt.Println("密码: ", string(pass))
	// Do something with pass

	// Silent. For printing *'s use gopass.GetPasswdMasked()
	fmt.Print("Enter [exit] go exit: ")
	input := bufio.NewScanner(os.Stdin)

	// 输入扫描
	for input.Scan() {
		text := input.Text()
		if text == "exit" {
			break
		} else {
			showPswd()
		}
		break
	}
}
