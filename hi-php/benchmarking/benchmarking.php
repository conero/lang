<?php
// 测试工具
function getCurTimes(){
    list($usec, $sec) = explode(" ", microtime());
    return ((float)$usec + (float)$sec);
}

//获取时间回调
function getTimeUsageCall(){
    $star = getCurTimes();
    return function() use($star){
        return getCurTimes() - $star;
    };
}


// 内存消耗计算
function getMemoryUsageCall(){
    $star = memory_get_usage();
    return function() use($star){
        return memory_get_usage() - $star;
    };
}
