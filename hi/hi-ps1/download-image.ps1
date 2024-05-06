# Joshua Conero
# 2020年3月18日 星期三

# https://thispersondoesnotexist.com/image 下载ai生成图片
$url = "https://thispersondoesnotexist.com/image"

#invoke-webrequest https://thispersondoesnotexist.com/image -outfile fake-people.jpg


# 循环生成
function save-image($count=100, $name=''){
    echo "将生产图片数: $count. "

    # return 'yes'
    # 日期
    $today=Get-Date
    $today = $today.ToString('yyyyMMdd')

    $path = './thispersondoesnotexist'
    if(-not (Test-Path -Path $path )){
        mkdir $path
    }

    for($i=1;$i -le $count;$i++)
    {
        invoke-webrequest $url -outfile "$path/$today-$i.jpg"
    }
}

save-image 300