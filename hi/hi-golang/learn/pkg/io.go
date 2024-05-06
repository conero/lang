package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
)

type InsWt struct {
	io.Writer
	Content string
}

func (iw *InsWt) Write(p []byte) (n int, err error) {
	content := string(p)
	iw.Content = content
	fmt.Println("内部程序写入: " + content)
	return 0, nil
}

type TestCase struct {
}

func (t TestCase) Writer() {
	tpl := `
		模板参数：  {{.Surong}}
`
	ins, err := template.New("t1").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	wt := &InsWt{}
	data := map[any]any{
		"Surong": 8 * 9,
	}
	ins.Execute(wt, data)

	fmt.Print(wt.Content, 5)
}

func main() {
	var t TestCase
	t.Writer()
}
