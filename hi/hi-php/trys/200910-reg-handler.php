<?php
// 解析 js 代码段：如[ut=(a(716),a(350)),At=a(309),dt=a.n(At),bt=a(310),gt=a.n(bt),ht=a(311),ft=a.n(ht),Mt=a(312),Et=a.n(Mt),vt=(a(718),a(313)),xt=a.n(vt),yt=a(314),zt=a.n(yt),jt=a(315),]
// 以","分个
function parseJs($js){
    $ss = $js;
    $reg = '/\([^\(\)]+\)/';
    $count = 1;
    $cacheMap = [];
    $maxCount = 5000;
    while(true){
        preg_match_all($reg, $ss, $match);
        $match = $match[0] ?? [];
        if (empty($match) || $count > $maxCount) break;
        
        for($i=0; $i<count($match); $i++){
            $vs = $match[$i];
            $vkey = '_$'.$count.'_';
            $cacheMap[$vkey] = $vs;

            $ss = str_replace($vs, $vkey, $ss);
            $count += 1;
        }

    }

    $lines = explode(',', $ss);

    $count = 1;
    while(true){
        if($count > $maxCount){
            break;
        }
        $count += 1;
        $delKey = [];

        $isMatch = false;
        foreach($lines as $i => $ln){
            foreach($cacheMap as $k => $vs){
                if(strpos($ln, $k) !== false){
                    $ln = str_replace($k, $vs, $ln);
                    $lines[$i] = $ln;
                    $delKey[] = $k;
                    $isMatch = true;
                }
            }
        }

        if(false === $isMatch) break;
    }

    return $lines;
}


// 代码测试
function test(){
    $js = 'ut=(a(716),a(350)),At=a(309),dt=a.n(At),bt=a(310),gt=a.n(bt),ht=a(311),ft=a.n(ht),Mt=a(312),Et=a.n(Mt),vt=(a(718),a(313)),xt=a.n(vt),yt=a(314),zt=a.n(yt),jt=a(315),';
    $lines = parseJs($js);
    print_r($lines);
}
// console
test();