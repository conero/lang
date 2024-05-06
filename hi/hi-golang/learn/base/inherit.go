package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

// 2018年8月3日 星期五
// 项目集成测试
// 简单MVC 模型

// 模式类型
type Model struct {
	Data map[any]any
}

func (m *Model) Get(key any) any {
	v, has := m.Data[key]
	if !has {
		v = nil
	}
	return v
}

// 设置值
func (m *Model) Set(key, value any) *Model {
	if nil == m.Data {
		m.Data = map[any]any{}
	}
	m.Data[key] = value
	return m
}

// 视图类型
type View struct {
	DefKey string
	Vid    string
	Vdata  map[string]string
}

func (v *View) GetVdfk() string {
	vk := "default"
	if "" != v.DefKey {
		vk = v.DefKey
	} else {
		v.DefKey = vk
	}
	return vk
}

// 视图处理
func (v *View) AddStr(str string) *View {
	if nil == v.Vdata {
		v.Vdata = map[string]string{}
	}
	s, h := v.Vdata[v.GetVdfk()]
	if h {
		s += str
	} else {
		s = str
	}
	v.Vdata[v.GetVdfk()] = s
	return v
}

// v1 = id, v2 = Data
func (v *View) GetVString(a ...any) {
	key := v.GetVdfk()
	if len(a) > 0 {
		if inStr1 := a[0].(string); inStr1 != "" {
			key = inStr1
		}
	}
	data := map[any]any{}
	if len(a) > 1 {
		data = a[1].(map[any]any)
	}
	if v.Vid == "" {
		v.Vid = "tDefault"
	}
	if t1, has := v.Vdata[key]; has && t1 != "" {
		tpl, err := template.New(v.Vid).Parse(t1)
		if err != nil {
			log.Fatal(err.Error())
		}
		err2 := tpl.Execute(os.Stdout, data)
		if err2 != nil {
			log.Fatal(err2.Error())
		}
	}

}

// 控制器
type Controller struct {
	Events  map[string]any
	InitKey string
	EndKey  string
}

func (c *Controller) _call(Name string) bool {
	if fn, has := c.Events[Name]; has {
		fn.(func())()
		return true
	}
	return false
}
func (c *Controller) Run(Name string) {
	// Init,
	if c.InitKey == "" {
		c.InitKey = "Init"
	}
	c._call(c.InitKey)
	c._call(Name)
	// End,
	if c.EndKey == "" {
		c.EndKey = "End"
	}
	c._call(c.EndKey)
}

func (c *Controller) On(Name string, fn any) *Controller {
	if nil == c.Events {
		c.Events = map[string]any{}
	}
	c.Events[Name] = fn
	return c
}

// Mvc 接口
type MvcInter interface {
	// 输出内容
	Name() string
}

// mvc 结构体
type Mvc struct {
	Model
	View
	Controller
	MvcInter
}

func (m *Mvc) ShowView() {
	m.GetVString("", m.Data)
}

// 名称
func (m Mvc) Name() string {
	return "MVC"
}

// 视图层
func NewMvc() *Mvc {
	mvc := Mvc{}
	return &mvc
}

type TestCase struct {
}

// 基础测试
func (tc TestCase) Base() {
	mvc := NewMvc()
	rId := "Base"
	mvc.On(rId, func() {
		fmt.Println(1)
		mvc.AddStr(`EORR:  {{.Test}}`)
		mvc.AddStr("\r\n")
		mvc.AddStr(`.    :)   {{.surong}}`)

		mvc.Set("Test", "emma-test")
		mvc.Set("surong", 5*8)
		mvc.ShowView()
	})
	mvc.Run(rId)
}

func main() {
	var tc TestCase
	tc.Base()
}
