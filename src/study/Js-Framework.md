# Js-Framework
- 2017年11月8日 星期三

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

