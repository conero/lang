<?php
//循环测试

require_once(__DIR__.'/benchmarking.php');

function TestLoop($count, $calls=null){
    $time1 = getTimeUsageCall();
    $memory = getMemoryUsageCall();
    $data = [];
    for($i=0; $i<$count; $i++){
        $data[] = $i%2+$i;
        if(is_callable($calls)){
            $calls($i);
        }
    }
    return [
        'count' => $count,
        'memory_get_usage' => $memory(),
        'memory_get_peak_usage' => memory_get_peak_usage(),
        'time_sec' => $time1(),
    ];
}


//千万级别
$count =10000*1000;
$ret = TestLoop($count);
print_r($ret);

//千万级别
$count =10000*10000;
$ret = TestLoop($count);
print_r($ret);