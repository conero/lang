# css (Cascading Style Sheets)

> 2018年7月25日 星期三
>
> Joshua Conero



## 概述

- 发明时间： 1994
- 作者： 哈坤·利提  (Håkon Wium Lie)



> 互联网方发明者

- Tim Berners-Lee 
- Robert Cailliau



## 变量单位

- 



## Scss/Sass

>  scss  sass3 引入新语法，完全兼容 css3； 属于css的扩展

- 与sass的差异		SCSS 需要使用分号和花括号而不是换行和缩进 



### 特性

1. 可使用变量
2. 嵌套CSS

   1. 父选择器的标识符&
   2. 群组选择器的嵌套
   3. 嵌套属性
3. 导入SASS文件
4. 静默注释
5. 混合器
6. 使用选择器继承来精简CSS



> 变量的使用

```scss
// 申明变量
$height: 500;
p{
   height: $height;
}
h1{margin: 0 0 $height 0;}

```



> 嵌套css

```scss
// 1
#main {
    p {
        a {
            color: red;
        }
    }
    .name{margin:0;}
    h1{font-size: 2px;}
}

// 编译为 css
#main p a{color: red;}
#main .name{margin:0;}
#main h1{font-size: 2px;}


// 2
// 属性嵌套
nav {
  border: {
  style: solid;
  width: 1px;
  color: #ccc;
  }
}
// => 
nav {
  border-style: solid;
  border-width: 1px;
  border-color: #ccc;
}
```



> 父选择器的标识符&

```scs
#content aside {
  color: red;
  body.ie & { color: green }
}

/*编译后*/
#content aside {color: red};
body.ie #content aside { color: green }
```



> 群组选择器的嵌套

```scss
nav, aside {
  a {color: blue}
}

/*编译后*/
nav a, aside a {color: blue}
```



> 样式继承

```scs
//通过选择器继承继承样式
.error {
  border: 1px solid red;
  background-color: #fdd;
}
.seriousError {
  @extend .error;
  border-width: 3px;
}
```

