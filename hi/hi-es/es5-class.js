/**
 * ES5 类，实现
 */


/**
 * 基础用户类
 */
function BasePerson(){
    this.name = null;
    this.sex = 1;
    this.birthday = null;

    // 私有变量
    var privateKey = Math.random();

    // 获取随机数
    this.myrand = function(){
        return privateKey;
    }

    // 获取数据id
    this.getId = function(){
        var temp = privateKey + '';
        var queque = temp.split('.');
        return queque[1];
    }
}


/**
 * prototype 式继承
 */
BasePerson.prototype = {
    getName: function(){
        return this.name;
    },
    setName: function(name){
        return this.name = name;
    }
};




/**
 * 继承类
 */
function JoshuaConero(){
    BasePerson.call(this);
    
    this.name = 'Joshua Conero';

    // 重写
    this.getName = function(){
        return 'Joshua Conero';
    }

    this.getAddress = function(){
        return 'Guiyang';
    }
}

JoshuaConero.prototype = new BasePerson();
JoshuaConero.prototype.constructor = JoshuaConero;
// 覆盖继承类的: prototype
// JoshuaConero.prototype.getAddress = function(){
//     return 'Guiyang';
// }

(() => {
    var jc = new JoshuaConero();

    jc.name = 'Joshua Uymas Conero';

    console.log(jc.getName());
    console.log(jc.setName(jc.getId()), jc.name);
    console.log(jc.getAddress());

    return jc;
})();


// ----------------------[测试语句]--------------------------
/*

// 调用1
// BasePerson 类的使用
(() => {
    var p = new BasePerson();
    p.name = 'Joshua Conero';

    console.log(p.getName());
    console.log(p.setName(p.getId()), p.name);

    return p;
})();


// 调用1
// JoshuaConero 继承效果

*/
