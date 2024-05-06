/**
 * 2017年11月14日 星期二
 * 排序算法参考
 */

/**

    // 伪代码
    input: an array a of length n with array elements numbered 0 to n − 1
    　　inc ← round(n/2)
    　　while inc > 0 do:
    　　for i = inc .. n − 1 do:
    　　temp ← a[i]
    　　j ← i
    　　while j ≥ inc and a[j − inc] > temp do:
    　　a[j] ← a[j − inc]
    　　j ← j − inc
    　　a[j] ← temp
    　　inc ← round(inc / 2.2)
 */


function shellSort(){
    var arr=[49,38,65,97,76,13,27,49,55,04],
        len =  arr.length;
    for(var fraction = Math.floor(len/2); fraction>0; fraction = Math.floor(fraction/2)){
        for(var i=fraction; i<len; i++){
            for(var j=i-fraction;
                j>=0 && arr[j]>arr[fraction+j] ; j-=fraction){
                var temp=arr[j];
                arr[j] = arr[fraction+j];
                arr[fraction+j] = temp;
            }
        }
    }
    console.log(arr);
}
