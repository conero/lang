/**
 * @date 2017年10月30日 星期一
 * promise 语法的实现
 */

// resolve, reject
// Promise 自定义实现
function Prms(callback){
} 


function TestPromise(isE){
    try {        
    } catch (error) {        
    }
    return new Promise(function(resolve, reject){
        if(isE){
            reject();
        }else{
            resolve();
        }
    });
}



(function(testError){
    // Promise 测试
    new Promise(function(t, v){ 
        /*       
        var isE = t();
        console.log(isE);
        if(isE){
            v();
        }    
        */
        if(testError){
            v();
        }else{
            t();
        }
    }).then(function(){
        console.log(Math.random());
        return true;
    }).catch(function(){
        console.log('Bad');
    });
})();