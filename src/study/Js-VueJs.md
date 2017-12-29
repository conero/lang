# VueJs(学习笔记)
- 20170717
> 
    author: Evan You
    the progressive javascript framework (渐进式javascript框架)

## 简介

    Vue 采用自底向上增量开发的设计
    -> 视图层
    -> 能为复杂的单页应用程序提供驱动
    -> 虚拟 DOM 原理
    -> 类 MVVM 框架
    -> 20kB min+gzip 运行大小

## 笔记    
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

## 组件组合
- 在 Vue 中，父子组件的关系可以总结为 *** prop 向下传递，事件向上传递 ***。父组件通过 prop 给子组件下发数据，子组件通过事件给父组件发送消息。
    - ![vue-lifecycle](./props-events.png)