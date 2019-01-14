# JS
- 2017年11月8日 星期三



## 工具

### npm

**主要命令介绍**

[参照](https://docs.npmjs.com/cli-documentation/cli)

```powershell
# 根据默认初始化库
npm init -y
# 交互式初始化
npm init

# npm install <tarball url>
# https://github.com/<user>/<project>.git#<branch>
#   [options]
#	--save-dev     已开发的模式保存
```



### [jest](https://jestjs.io)

*js 测试框架*

> https://github.com/facebook/jest



## Browser



### MouseEvent

*坐标点：* ![](./image/js/x-y.png)

```js
clientX, clientY    // 鼠标指针向对于浏览器页面（当前窗口）的坐标

layerX, layerY		//

movementX, movementY

offsetX,offsetY

pageX, pageY		// 文档坐标而非窗口坐标,这2个属性不是标准属性，但得到了广泛支持。IE事件中没有这2个属性

x, y
    
screenX, screenY
```





## MVVM



## jQuery



## ReactJs

*React is a declarative, efficient, and flexible JavaScript library for building user interfaces. It lets you compose complex UIs from small and isolated pieces of code called “components”.*



- A JavaScript library for building user interfaces (*用于构建用户界面的JavaScript库*)
    - facebook
- jsx   javascript xml， XML-like syntax called JSX 
- 自定义类型，必须以大写开头，否则将视作 html 文档格式
- React.Component 组件
    - this.props 传入的属性值 如： <React.Component title="85"> => this.props.title == '85'
    - this.props 属性，通过 ``render()`` 函数渲染
    - this.state 状态值
      - this.setstate 设置状态值



### Component 组件

*组件接受称为属性(`props` properties 的简称)的参数，通过方法`render` 渲染为一个有层次化(结构)的视图。*

- 函数式 ``const element = () => <h1>标题</h1>``

  ```jsx
  function Welcome(props) {
    return <h1>Hello, {props.name}</h1>;
  }
  const element = <Welcome name="Sara" />;
  ReactDOM.render(
    element,
    document.getElementById('root')
  );
  ```

- 类式   ``extends React.Component``









## AngularJS



## VueJs

- 20170717
>
>   author: Evan You
>   the progressive javascript framework (渐进式javascript框架)

### 简介

    Vue 采用自底向上增量开发的设计
    -> 视图层
    -> 能为复杂的单页应用程序提供驱动
    -> 虚拟 DOM 原理
    -> 类 MVVM 框架
    -> 20kB min+gzip 运行大小

### 笔记    
- 字符串模板语法， 支持： mustache 模板语法
- HTML 原始属性，通常以 ***v-*** 打头，也有相应的简写方法
    - v-bind    数据绑定
        - v-bind:id = "idx"     => :id = "idx"
        - v-bind:href = "xf"    => :href = "xf"
    - v-html    原始 HTML
    - v-on      事件绑定
        - v-on:click    => @click
    - v-model   表单语法, 表单控件元素上创建双向数据绑定
    - v-if/v-for 条件等语句判断    
- template 组件，通过浏览器编译 js, ***.vue*** 视图文件
    - 组件组合:     prop 向下传递，事件向上传递
    - 注册全局组件:   Vue.component(tagName, options)
    - 局部组件:
```javascript
    var Child = {
        template: '<div>A custom component!</div>'
    }
    new Vue({
    // ...
    components: {
        // <my-component> 将只在父组件模板中可用
        'my-component': Child
        }
    })
```
- vue 实例生成周期
    - ![vue-lifecycle](./Js-VueJs-lifecycle.png)

### 组件组合
- 在 Vue 中，父子组件的关系可以总结为 *** prop 向下传递，事件向上传递 ***。父组件通过 prop 给子组件下发数据，子组件通过事件给父组件发送消息。
    - ![vue-lifecycle](./props-events.png)







# _WebGL<small>(Web Graphics Library)</small>_

常见图形化库：

- [three.js](https://threejs.org/)
- [p5.js](https://p5js.org/)
- [D3.js](https://d3js.org/)



# NodeJs



## 收藏

- [*node-ffi nodejs 调用 dll库*](https://www.npmjs.com/package/ffi)

