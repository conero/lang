# Joshua Conero
# 2021年10月21日 星期四
# 多线程并发请求网站。

$session = New-Object Microsoft.PowerShell.Commands.WebRequestSession
$session.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36"
$session.Cookies.Add((New-Object System.Net.Cookie("PHPSESSID", "967b09efb2f771edec215102c209d500", "/", "www.gy-imcloud.com")))

for($i=0; $i -lt 100000; $i++){
    Write-Output $(Get-date -Format "HH:mm:ss") "请求： $i";
    start-job -ScriptBlock {
        Invoke-WebRequest -UseBasicParsing -Uri "https://www.gy-imcloud.com/index/video/lists" `
            -Method "POST" `
            -WebSession $session `
            -Headers @{
            "sec-ch-ua"="`"Google Chrome`";v=`"95`", `"Chromium`";v=`"95`", `";Not A Brand`";v=`"99`""
            "Accept"="*/*"
            "X-Requested-With"="XMLHttpRequest"
            "sec-ch-ua-mobile"="?0"
            "sec-ch-ua-platform"="`"Windows`""
            "Origin"="https://www.gy-imcloud.com"
            "Sec-Fetch-Site"="same-origin"
            "Sec-Fetch-Mode"="cors"
            "Sec-Fetch-Dest"="empty"
            "Referer"="https://www.gy-imcloud.com/index/index/video_list"
            "Accept-Encoding"="gzip, deflate, br"
            "Accept-Language"="zh,en-GB;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6"
            } `
            -ContentType "application/x-www-form-urlencoded; charset=UTF-8" `
            -Body "keyword=%E5%B7%A5%E6%8A%95"
    }
}
