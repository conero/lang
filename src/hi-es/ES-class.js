/**
 * Es Class 测试以及学习
 * 2017年9月11日 星期一
 * https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Classes
 */

 // 普通类
class A{
    constructor(){
        this.createId = Math.random()
    }
    // 本身实例创建
    static instance(){
        return new this
    }
}


class Test{
    constructor(){
        this.AtestCase()
    }
    AtestCase(){
        var a = A.instance()
        console.log(a)
    }
}

new Test