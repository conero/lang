# PowerShell

> 2018年10月19日 星期五




## 简介

*Windows PowerShell® 是基于任务的命令行管理程序和脚本语言，专为进行系统管理而设计。*

*UNIX 系统一直有着功能强大的脚本（shell），Windows PowerShell 的诞生就是要提供功能相当于 UNIX 系统 BASH 的命令行外壳程序，同时也内建脚本语言以及辅助脚本程序的工具。*

*PowerShell 使用“谓词 - 名词”命名系统。 每个 cmdlet 名称都由一个标准谓词、连字符和特定名词组成。*

*PowerShell 是构建于 .NET 上基于任务的命令行 shell 和脚本语言。*



> 设计目标

*PowerShell 旨在消除长期存在的问题和添加新功能，从而改进命令行和脚本环境。*



*Windows系统下 shell 脚本无效，windows系统用文件名的后缀来判断文件类型，如只要是.py后缀的就关联到python程序去执行。而Linux之类通过文件首行标记来决定，因此如下脚本无效：*

```shell
#!/user/bin/env python
```
或者
```shell
#!/user/bin/python
```

而采用

```powershell
>> ./run-python.py
```





> 基本规则

1. 命令等不区分大小写
2. 面向对象（结构化信息）
3. PowerShell 基于 .NET Framework 构建。 它与 C# 编程语言共享一些语法功能和关键字
4. 使用 PowerShell 命令，参数名称前面始终带有“-”
5. 星号 (*) 用于 PowerShell 命令参数中的通配符匹配。 * 表示“匹配一个或多个任意字符”。 
6.  使用 PowerShell，可通过按 Tab 键来填充文件名和 cmdlet 名称
7. 默认情况下，在处理文本时，PowerShell 比较运算符不区分大小写。





## 变量

*PowerShell 处理对象。 使用 PowerShell 可以创建称为“变量”的命名对象。 变量名称可以包含下划线字符，也可以是任何字母数字字符。 在 PowerShell 中使用时，始终使用变量名称后跟的 $ 字符指定变量。*



```ini
# 赋值
$jc = 'Joshua Conero is a good guy。'

$jc = $jc + 'We all should be nice to him.'

# Joshua Conero is a good guy。We all should be nice to him.
$jc

# 作为其它变量的别名
$jc = ls
$jc

# 获取环境变量
$jc = Get-ChildItem env:
$jc
# 读取环境变量
$env:path
# 设置环境变量
$env:path = "new value"

# 查看所有变量
Get-Variable

# 空判断，`$null` 默认的系统值，可用于判断控制。任务未定义的值都等于`$null`
if($null -eq $value){}
# 判断是否为有效值
if(($null -eq $value) -or ($value.Length -eq 0)){}
```



> *转义字符*

- "`"   字符可转移变量，原样输出; 也换行输入代码
- `'`   单引号内的变量不会成为模板
- `"`   双引号可为变量

```powershell
$n = 8 * 9

echo '$n is $n.'
# $n is $n.

echo "`$n is $n."
# $n is 72.
```



> 指示特殊字符

|          |                                                              |                                                              |
| -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 特殊字符 | 说明                                                         | 示例                                                         |
| `0       | Null。Powershell可以识别null特殊字符(`0),并用字符代码0来表示该特殊字符，Powershell输出中显示为空白。所以Powershell可以读取并处理包含null字符(字符串终止或记录终止指示符)的文本文件。需要注意，null 特殊字符与 $null 变量不同，后者存储 NULL 值。 |                                                              |
| `a       | 警报。可向计算机的扬声器发送蜂鸣信号，可用此字符向用户发出有关危险操作的警告 | PS C:\> for ( $i =0 ; $i -le 1; $i++) { "`a" }               |
| `b       | 退格。将光标后退一个字符。                                   | PS C:\> "Windows`b Powershell"Window Powershell              |
| `f       | 换页。打印字符，指示在当前字符的下一页继续打印，该字符只影响打印的文档，不影响屏幕输出 |                                                              |
| `n       | 换行。其后的内容在下一行显示。                               | PS C:\> "This line has been breaked into`n two lines"This line has been breaked into two lines |
| `r       | 回车符。会删除该字符之前的整行内容                           | PS C:\> Write-Host "will be deleted `r others" others        |
| `t       | 水平制表符。默认情况下，Powershell每隔7个空格为一个制表符。  | PS C:\> "columnA`tcolumnB`tcolumnC"columnA columnB columnC   |
| `v       | 垂直制表符。 光标前进到下一个垂直制表位并从该处开始写入后面的所有输出。该字符仅影响打印的文档，不影响屏幕输出。 |                                                              |



> 行号符号

```ini
`			行号符号
```



> Object 对象

```powershell
# 执行
ls
Get-ChildItem

# 其接口类似与: [{Mode:x, LastWriteTime: x, Name:x}]
<#
Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d----           2020/5/14    11:10                .idea
d----            2020/5/6    22:17                assets
d----           2020/5/13    14:02                bin
d----           2020/5/13    16:15                config
d----            2020/5/6     9:58                dist
d----           2020/5/14    10:44                doc
d----           2020/5/12    20:25                node_modules
d----          2019/10/27    10:26                public
d----            2020/3/3    11:08                runtime
-a---           2020/2/23    22:55            686 .gitignore
-a---           2020/2/23    23:30            820 cmd_db.go
-a---            2020/5/6    10:35           2657 distribute.ps1
-a---           2020/4/12    21:22            360 distribute-pack.ps1
-a---           2020/5/13     9:21           1179 go.mod
-a---           2020/5/13     9:21          51435 go.sum
-a---           2020/5/12    20:25           1487 package.json
-a---           2020/2/26    16:31          28799 package-lock.json
-a---            2019/9/3     9:51             61 README.en.md
-a---           2020/4/12    21:22           2423 README.md
-a---           2020/5/13    16:00           3241 yangsu.go
-a---           2020/5/12    20:25         223707 yarn.lock
#>

```





### 常量

```powershell
# ps脚本所在的目录
echo $PSScriptRoot

#
# 运行制定脚本内的脚本
./path/file.ps1
# >>>>>
# 文件： file.ps1
&"$PSScriptRoot/grafana-server.exe"
```



#### 命令行参数 `$args`

```powershell
# 判断
if ($args -contains "-d"){
	echo "获取的参数: -d"
}

# 数组包含运算
# -contains    	数组中包含元素
# -ccontains   	数组中包含元素(大小写敏感)
# -notcontains 	数组中不包含元素
# -cnotcontains	数组中不包含元素(大小写敏感)
```







### 对象/结构

*定义数组/结构等*

```powershell
# 数组
$arr = @(20, 20, 5, 14)
# 等同于: 
20,20,5,14


# 结构化
# name => value
$obj = @{
	surname = 'conero'
	first_name = 'Joshua'
    gender = 'A'; random = get-random
    age=18
}

# 输出当前所在目录
$tmp = Get-Location; $tmp.path
```



>  *`Get-Member` 查看对象结构*

```powershell
# 若要查看进程对象的所有成员并分页显示输出，以便于你可以全部查看
Get-Process | Get-Member | Out-Host -Paging
# 筛选类型； -MemberType 处理可以设置未显示的至外还可有> `All` 等
Get-Process | Get-Member -MemberType Property

```



> `Select-Object`  选择对象部件

```powershell
# Get-ChildItem 默认显示: Mode LastWriteTime Length Name

Get-ChildItem | Select-Object -Property length,name
#>> Length Name
#>> ------ ----

# 查看磁盘可用的大小尺寸
Get-CimInstance -Class Win32_LogicalDisk | Select-Object -Property Name,FreeSpace
# 改为 GB 单位
Get-CimInstance -Class Win32_LogicalDisk |
  Select-Object -Property Name, @{
    label='FreeSpace'
    expression={($_.FreeSpace/1GB).ToString('F2')}
  }
```



> `Where-Object` 过滤

```powershell
Get-CimInstance -Class Win32_SystemDriver |
  Where-Object {$_.State -eq "Running"} |
    Where-Object {$_.StartMode -eq "Auto"}
```



> 其他

```powershell
# 
# 排序
# Sort-Object
Get-ChildItem |
  Sort-Object -Property LastWriteTime, Name |
  Format-Table -Property LastWriteTime, Name
  
# 逆向排序   -Descending
Get-ChildItem | Sort-Object -Property Name -Descending


# 
# 遍历: ForEach-Object
@(1,3,5,6,7,9,100) | ForEach-Object -Process {$_*10}

#
# New-Object 创建 .Net Framework 对象/ COM 

# 使用静态方法
# 引用静态的 System.Environment 类
[System.Environment]
# 查看静态类名称
[System.Environment] | Get-Member -Static
[System.Math] | Get-Member -Static -MemberType Methods
```



> **WMI**

*Windows Management Instrumentation (WMI) 是 Windows 系统管理的核心技术，因为它以统一的方式公开大量信息。*



```powershell
# 查看 WMI 类列表
Get-CimClass -Namespace root/CIMV2 |
  Where-Object CimClassName -like Win32* |
    Select-Object CimClassName
# 查询 WMI 类，并保存到 list.json 文件中
Get-CimClass -Namespace root/CIMV2 |
	Where-Object CimClassName -like Win32* |
	Select-Object CimClassName | ConvertTo-Json > list.json


# 显示 WMI 类详细信息
Get-CimInstance -Class Win32_OperatingSystem
```



### 运算符

*脚本块使用特殊变量 `$_` 来指代管道中的当前对象。*

| 比较运算符   | 含义                       | 示例（返回 True）            |
| :----------- | :------------------------- | :--------------------------- |
| -eq          | 等于                       | 1 -eq 1                      |
| -ne          | 不等于                     | 1 -ne 2                      |
| -lt          | 小于                       | 1 -lt 2                      |
| -le          | 小于或等于                 | 1 -le 2                      |
| -gt          | 大于                       | 2 -gt 1                      |
| -ge          | 大于或等于                 | 2 -ge 1                      |
| -like        | 相似（文本的通配符比较）   | "file.doc" -like "f*.do?"    |
| -notlike     | 不相似（文本的通配符比较） | "file.doc" -notlike "p*.doc" |
| -contains    | 包含                       | 1,2,3 -contains 1            |
| -notcontains | 不包含                     | 1,2,3 -notcontains 4         |



```powershell
1,2,3,4 | Where-Object {$_ -lt 3}
```



> 逻辑运算符

| 逻辑运算符 | 含义                                        | 示例（返回 True）        |
| :--------- | :------------------------------------------ | :----------------------- |
| -and       | Logical and；如果两侧都为 True，则返回 True | (1 -eq 1) -and (2 -eq 2) |
| -or        | Logical or；如果某一侧为 True，则返回 True  | (1 -eq 1) -or (1 -eq 2)  |
| -not       | Logical not；反转 True 和 False             | -not (1 -eq 2)           |
| !          | Logical not；反转 True 和 False             | !(1 -eq 2)               |



### 输入输出

> 调用 VB 做图形化输入

```powershell
# pwsh7 执行错误
[void][System.Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
$cn = [Microsoft.VisualBasic.interaction]::inputbox('Enter Your Name', 'Name', 'Jc')
echo ">>:  $cn"
```



> 命令输入/输出

```powershell
read-host "输入提示"
#>> 输入提示:

# 输出，感觉与 echo 一样. echo 应该更加强大吧
write-host "输出内容"
echo "输出内容"

# 指定颜色 (-ForegroundColor)
Write-Host "2021/10/27" -ForegroundColor Blue
```



> 输入格式处理

1. *Out-Host*   直接将数据发送到控制台
2. *Out-Null*   空输出
3. *Out-File*    文件保存
4. ConvertTo-Json     输出文本内容未 json 格式
5. ConvertTo-Csv      输出文本内容未 csv 格式

```powershell
# 列表输出
Get-Process | Out-Host -Paging | Format-List

# Out-Null 空输出
Get-Command | Out-Null

# 输出到文件
Get-Process | Out-File -FilePath ./processlist.txt
Get-Process > ./processlist2.txt

# 查看输出格式
# Verb 谓词
Get-Command -Verb Format -Module Microsoft.PowerShell.Utility
# 输出内容指定格式
Get-Command -Verb ConvertTo
```



### 类型

#### 字符串

> 字符串分割

```powershell
# 环境变量分割
$env:Path -split ";"

# 字符串包含
"123".contains("12")

#字符串数字连接
"Joshua","Jhan","Conero" -join "."

# 子字符串
"Joshua Conero".Substring(3,3)

# 字符串格式化
"{0} {2} {1}" -f "you","die","are"

# 字符串替换
"Joshua Lv U".Replace("Lv", "heart")

# 长度: >> 6	
"Joshua".Length

# 字符串分割
"Joshua".Split("s")
```





### 脚本语法

#### 算数运算

*以下为特殊的符号，其他:*

```powershell
# `+` 加

# 数字算术运算
2000+20
 
# 字符串/array/object 合并
'Joshua'+",Conero"
@(1, 3) + @('a', 'c')
@{'name'='Joshua'} + @{'age'= 28; 'last'='Coenro'}
```



#### 异常

> ErrorAction

`-ErrorAction`  可以在单条语句中简单的控制异常信息。别名`ea`，可设置的值 `Suspend, Ignore, Inquire, Continue, Stop, SilentlyContinue` 。参考[通用选项](https://docs.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/about/about_commonparameters)





## 命令参考

**cmdlet**

*Windows PowerShell 引入了 cmdlet（读作“command-let”）的概念，它是内置于 Shell 的简单的单一函数命令行工具。PowerShell 中的本机命令称为 cmdlet （读作 command-let）。*



```powershell
# 获取命令列表所有）
Get-Command
# 获取当前会话的所有命令集
Get-Command *

# get-command + tab键 可以快速查询
# 查找“Get-”开头的命令
get-command Get-*

# 查看命令已经分配的别名
Get-Command -CommandType Alias

#
# 查看帮助信息, [Get-Help] 与 [-?] 相同
Get-Help Write-Host
Write-Host -?
# man/help 雷士
man write-host

# 常用命名
# 打印当前 ps 的版本信息
$psversiontable
$host
```



*常用参数* (`cmdlet -param`)

```powershell
# 
#参数
#通用参数
# WhatIf, Confirm, Verbose, Debug, Warn, ErrorAction, ErrorVariable, OutVariable, OutBuffer

# 其他重要的建议参数名称是
# Force, Exclude, Include, PassThru, Path, CaseSensitive
# Exclude 排除
```





**单行执行多条代码**

```powershell
# 转入且初始化 npm
cd ./jc_cj | npm init -y
```





> 关机或重启

```ini
stop-computer
restart-computer
```



### 与 Cmd/ Unix 类似

可在 PowerShell 中使用的常见 cmd.exe 和 UNIX 命令：

|       |         |       |       |
| ----- | ------- | ----- | ----- |
| cat   | dir     | mount | rm    |
| cd    | echo    | move  | rmdir |
| chdir | erase   | popd  | sleep |
| clear | h       | ps    | sort  |
| cls   | history | pushd | tee   |
| copy  | kill    | pwd   | type  |
| del   | lp      | r     | write |
| diff  | ls      | ren   |       |

```powershell
# 删除多文件,忽视文件删除
rm -r ./path

# 查看别名,详情
Get-Alias cls
# 全部别名
Get-Alias 
gal
```



*PowerShell 别名尝试兼顾清晰度和简洁性。 PowerShell 为常见名词和谓词使用一组标准的别名。*

| 名词或谓词 | 缩写 |
| ---------- | ---- |
| Get        | g    |
| Set        | s    |
| Item       | i    |
| Location   | l    |
| Command    | cm   |
| Alias      | al   |





了解简写名称后，这些别名是可以理解的。

| Cmdlet 名称    | Alias |
| -------------- | ----- |
| `Get-Item`     | gi    |
| `Set-Item`     | si    |
| `Get-Location` | gl    |
| `Set-Location` | sl    |
| `Get-Command`  | gcm   |
| `Get-Alias`    | gal   |



### 管道

*管道的行为就像一系列连接的管道段一样。 沿着管道移动的项会通过每个管道段。 若要在 PowerShell 中创建管道，请使用管道运算符“|”将命令连接在一起。 每个命令的输出都将被用作下一命令的输入。*





### Start-Process

```powershell
Start-Process
     [-FilePath] <String>
     [[-ArgumentList] <String[]>]
     [-Credential <PSCredential>]
     [-WorkingDirectory <String>]
     [-LoadUserProfile]
     [-NoNewWindow]
     [-PassThru]
     [-RedirectStandardError <String>]
     [-RedirectStandardInput <String>]
     [-RedirectStandardOutput <String>]
     [-WindowStyle <ProcessWindowStyle>]	窗体样式设置
     [-Wait]		直到进行完成
     [-UseNewEnvironment]
     [<CommonParameters>]
```





用于本地进程开启的命令

```powershell
# 启动使用默认应用打开程序
# 开浏览器地址地址
Start-Process -FilePath http://localhost:2079/

# 使用id浏览器打开
# starts with a specific browser
Start-Process -FilePath iexplore -ArgumentList www.powershellmagazine.com

# 以管理员身份运行【powershell】
Start-Process -FilePath "powershell" -Verb RunAs -ArgumentList "file.ps1","argument-2"

# 本地打开应用
## -- powershell.exe
Start-Process powershell.exe -ArgumentList "-noexit"  
Start-Process powershell -ArgumentList "-noexit", "-noprofile", "-command &{Get-Location}"
## -- pwsh.exe
Start-Process pwsh.exe -ArgumentList "-noexit", "-noprofile", "-command &{Get-Location}"

# 启动 cmd，打开新窗口
Start-Process cmd.exe -ArgumentList "/C .\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties"
Start-Process cmd -ArgumentList "/C .\bin\windows\kafka-server-start.bat .\config\server.properties"
```



### Invoke-WebRequest

*使用方法: 在谷歌浏览器中打开任意web，开启`开发者工具`中的 `Network` 标签，依次选中其下`右键`的`Copy`，`Copy as PowerShell`等（`chrome/F12-开发者工具/Network/右键/Copy/Copy as PowerShell`），如下:*

```powershell
Invoke-WebRequest -Uri "http://127.0.0.1:9105/guest/member/index" `
-Method "POST" `
-Headers @{
"Accept"="application/json, text/javascript, */*; q=0.01"
  "X-Requested-With"="XMLHttpRequest"
  "User-Agent"="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"
  "Origin"="http://127.0.0.1:9105"
  "Sec-Fetch-Site"="same-origin"
  "Sec-Fetch-Mode"="cors"
  "Sec-Fetch-Dest"="empty"
  "Referer"="http://127.0.0.1:9105/guest/member/index"
  "Accept-Encoding"="gzip, deflate, br"
  "Accept-Language"="zh-CN,zh;q=0.9"
  "Cookie"="deviceid=1565750664628; xinhu_ca_adminuser=admin; xinhu_ca_rempass=0; xinhu_mo_adminid=rjr0rjq0ds0rjd0aa0dl0dd0rlj0rjr0dl0hn0nh09; lang_current=zh-hans; yangsu=e8ec7af7-f523-4fb5-9c73-e6dca6794a4d; lang=zh-CN; Hm_lvt_f1161c9c97f4c7b7231878dd77fe8b5a=1567136830,1567389984; mysession_cookie_name=9ca002ff-3c2f-4201-aeea-4245ab35564d; jenkins-timestamper-offset=-28800000; UM_distinctid=1702cb7aaad1d2-053921863d8eb8-b383f66-144000-1702cb7aaae26c; mycookiesessionnameid=e14d25d5-b9d7-4fdf-b11b-d90f0ba9ddb2; think_lang=zh-cn; from=none; PHPSESSID=ee7111ac5b206cb19e9a3ac7ccfe81dd; CNZZDATA1275935409=470955613-1564412279-%7C1589330253; sid=v7uzVpef8IIS8WkAdX6RCQ%3D%3D; ck=1019a40bd8e2731d34d692805b0cd735"
} `
-ContentType "application/x-www-form-urlencoded; charset=UTF-8" `
-Body "pageStart=1&pageSize=10"
-OutFile './test.data.json'

# 说明
# -OutFile 将请求的内容保存到文件中
```



### 文件/目录

*目录存在检测*

```powershell
$path = './runtime/bundle'
if(-not (Test-Path -Path $path )){
    mkdir $path
}

# 目录存在
Test-Path $dir -PathType Container
# 目录检测
Test-Path $filepath -PathType Leaf


# 路径
cd $path
```



*目录、文件查找*

```powershell
# 查找当前目录下以.txt结尾的文件
dir -name *.txt

# 指定目录下查找文件
# -Recurse 表示遍历子目录
dir -path ./*php -Recurse

# 别名:`dir -> Get-ChildItem`


# 查找文件并删除文件
Get-ChildItem -Recurse $path/config/*TAGprivate* | Remove-Item
```



*删除目录/文件*

```powershell
# 强制删除文件目录等
rm $path -recurse -force
```



**查找文件资源并删除它**

```powershell
# 相对目录需要修正路径
Get-ChildItem -Recurse -Name *ts | Remove-Item -Force

# 相对目录需要修正路径
Get-ChildItem .\uploads\videos\ -Recurse -Name *ts | ForEach-Object -Process {".\uploads\videos\$_"} | Remove-Item -Force
```



**查找按大小排序限定数文件**

```powershell
Get-ChildItem d:\ -Recurse | Sort-Object -Descending length | Select-Object -First 10
```





#### 文件下载

*图片批量下载*

```powershell
# https://thispersondoesnotexist.com/image 下载ai生成图片
$url = "https://thispersondoesnotexist.com/image"

#invoke-webrequest https://thispersondoesnotexist.com/image -outfile fake-people.jpg

# 循环生成
function save-image($count=100, $name=''){
    # 日期
    $today=Get-Date
    $today = $today.ToString('yyyyMMdd')

    for($i=1;$i -le $count;$i++)
    {
        invoke-webrequest $url -outfile "./thispersondoesnotexist/$today-$i.jpg"
    }
}

save-image
```



*批量文件下载*

```powershell
$vlist = @("https://v-cdn.zjol.com.cn/280443.mp4", "https://v-cdn.zjol.com.cn/276982.mp4", "https://v-cdn.zjol.com.cn/276984.mp4", "https://v-cdn.zjol.com.cn/276985.mp4", "https://v-cdn.zjol.com.cn/276986.mp4", "https://v-cdn.zjol.com.cn/276987.mp4", "https://v-cdn.zjol.com.cn/276988.mp4", "https://v-cdn.zjol.com.cn/276989.mp4", "https://v-cdn.zjol.com.cn/276990.mp4", "https://v-cdn.zjol.com.cn/276991.mp4", "https://v-cdn.zjol.com.cn/276992.mp4", "https://v-cdn.zjol.com.cn/276993.mp4", "https://v-cdn.zjol.com.cn/276994.mp4", "https://v-cdn.zjol.com.cn/276996.mp4", "https://v-cdn.zjol.com.cn/276998.mp4", "https://v-cdn.zjol.com.cn/277000.mp4", "https://v-cdn.zjol.com.cn/277001.mp4", "https://v-cdn.zjol.com.cn/277002.mp4", "https://v-cdn.zjol.com.cn/277003.mp4", "https://v-cdn.zjol.com.cn/277004.mp4")


$vlist | ForEach-Object -Process {
    $queue = $_ -split "/"
    $filename = $queue[$queue.Length-1]
    Invoke-WebRequest -Uri $_ -OutFile "./$filename"
}
```





### 计算机管理

> **WMI**

*Windows Management Instrumentation (WMI) 是 Windows 系统管理的核心技术，因为它以统一的方式公开大量信息。*



```powershell
# 查看 WMI 类列表
Get-CimClass -Namespace root/CIMV2 |
  Where-Object CimClassName -like Win32* |
    Select-Object CimClassName
# 查询 WMI 类，并保存到 list.json 文件中
Get-CimClass -Namespace root/CIMV2 |
	Where-Object CimClassName -like Win32* |
	Select-Object CimClassName | ConvertTo-Json > list.json


# 显示 WMI 类详细信息
Get-CimInstance -Class Win32_OperatingSystem
```





```powershell
# 本地计算机上系统 BIOS 的高度压缩的完整信息
Get-CimInstance -ClassName Win32_BIOS

# 获取本地完整的记录软件
Get-CimInstance -ClassName Win32_Product
```



**导出计算机中应用安装列表**

```powershell
Get-CimInstance -ClassName win32_product | ConvertTo-Json > win32_product.json
```





#### 进程/服务

- 进程    *Get-Process*
- 服务 
  - *Get-Service*             查询服务
  - *New-Service*            创建服务
  - *Remove-Service*       移除服务
  - Restart-Service
  - Start-Service
  - Stop-Service



```powershell
#
# 进程
# 列出所有进程，与 ps 相同
Get-Process

# 进程搜索，如下
ps -Name wech*

# 关闭进程
Stop-Process -Name $name
# 关闭 firefox
Stop-Process -name firefox


#
# 服务
# 查看，mysql 服务
Get-Service -name mysql*
# 获取从属服务
Get-Service -Name LanmanWorkstation -RequiredServices
# 获取依赖服务
Get-Service -Name LanmanWorkstation -DependentServices

# 启动/停止服务
# 通过`-name` 查询
Stop-Service 
start-service
# 重启
restart-service

# 服务
Start-Service Apache2.4

#
# 查询服务: `Apache2.4`
Get-CimInstance -ClassName Win32_Service -Filter "Name='Apache2.4'"
# 查询2
Get-Service Apache2.4
```



> `New-Service` 创建服务





**进程** **x-job**

```powershell
# 查询正在运行的进程
get-job -State Running
# 查询运行的进程并关闭
get-job -State Running|stop-job

# 关闭任务
stop-job -Id 1

# 启动任务（可实现并发的延迟执行）
start-job -ScriptBlock {
	sleep 5; Write-Host "Hello world.";
}

# 使用 ps1 脚本作为任务
start-job -name gzh-runing -FilePath ./script_mg.ps1
```







#### 驱动器

*Windows PowerShell 驱动器 是一个数据存储位置，你可以像访问 Windows PowerShell 中的文件系统驱动器那样访问它。 Windows PowerShell 提供程序将为你创建一些驱动器，例如文件系统驱动器（包括 C: 和 D:）、注册表驱动器（HKCU: 和 HKLM:）和证书驱动器 (Cert:)，你也可以创建自己的 Windows PowerShell 驱动器。 这些驱动器非常有用，但它们仅在 Windows PowerShell 内可用。 你无法通过使用其他 Windows 工具（如文件资源管理器或 Cmd.exe）访问它们。*



```powershell
# 查看驱动器
Get-PSDrive
# 文件驱动器
Get-PSDrive -PSProvider FileSystem
# 删除驱动器
Remove-PSDrive
```



> *设备查看*

```powershell
# 打印机
Get-CimInstance -Class Win32_Printer
```



#### WMI/CIM

```shell
# 通过SQL查询 Win32_BIOS 的信息。
Get-CimInstance -Query 'Select * from Win32_BIOS'

# 通过类名调用
Get-CimInstance -ClassName Win32_BIOS

#直接获取序列号
(Get-CimInstance -ClassName Win32_BIOS -Property SerialNumber).SerialNumber

# 搜索或查看类名
Get-WmiObject -List *
Get-WmiObject -List Win32_*
Get-WmiObject -List Win32_*|Select-Object name
```





### 其他

```powershell
# 获取 ip 地址
Get-NetIPAddress | Where-Object {$_.AddressFamily -eq 'IPv4'} | ForEach-Object IPAddress | ConvertTo-Json
```



*程序运行计时*

```powershell
$elapsed = [System.Diagnostics.Stopwatch]::StartNew() 
write-host "Started at $(get-date)" 
for ($t=1; $t -le 20; $t++) { 
    Write-Host "Elapsed Time: $($elapsed.Elapsed.ToString())" 
    sleep 1 
} 
write-host "Ended at $(get-date)" 
write-host "Total Elapsed Time: $($elapsed.Elapsed.ToString())" 
```



*环境变量*

> powershell 语言环境变量操作

```powershell
# 读取环境变量
ls env:
# ls 别名 get-childitem 
get-childitem env:
# 读取环境变量
echo $env:java_home
```



> cmd.exe 环境变量操作

```shell
# 输出环境变量
echo %JAVA_HOME%
# 设置环境变量
set Java_home= 'D:\Program Files\java\xxxxxxx'

# windows 下可使用 setx 命令
setx JC_HOME C:/test
# 注册表写入（管理模式下）
setx /M BAT_HOME_JC C:\bat
```

[comment]: <> "<!-- @todo 删除环境变量的方法 -->"
[//]: # "<!-- @todo 删除环境变量的方法 -->"

#### 日期/时间

```powershell
# 获取时间对应的格式数据
Get-Date -Format "dddd MM/dd/yyyy HH:mm K"
Get-Date -Format 'yyMMddHHmmss'
```







## CMD

*cmd 大部分命令在powershell上是兼容，只是其语法规则与其有差别。*

```shell
# 命令行帮助
mkdir /?

# cmd.exe 单行运行多条命令
mkdir .\jc_cj
cd .\jc_cj && npm init -y
rmdir /s .\jc_cj

# 单行处理多脚本
mkdir .\jc_cj && cd .\jc_cj && npm init -y && rmdir /s .\jc_cj
```



### where

>  查找项目/命令所在目录

```shell
# cmd
# 查询 svn 所在目录
where svn
where go

# powershell 中使用 where 命令
where.exe svn

# 查看以“net”开头的文件
where.exe net*

# dir 与 其他 ls 命令差不多
dir
```



### 服务/进程

- **service**                          服务

- **process**						 进程



> 删除服务

```powershell
# 删除 window 10 服务的口令
# cmd.exe 成功， powershell 无效
sc delete <name>
```



> 进程管理

```powershell
# [tasklist] 进程列表
# 进制管理
# 查看帮助
tasklist /?
tasklist /fi "imagename eq nginx.exe"

# [taskkill] 进制终止
taskkill /?
# 强制杀死进程
taskkill /fi "imagename eq nginx.exe" /F
```



### 系统

**systeminfo**

*该工具显示本地或远程机器(包括服务包级别)的操作系统配置的信息。*





## issue

### iwr

> iwr : 请求被中止: 未能创建 SSL/TLS 安全通道。

```powershell
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
```



*对应 C #*

```c#
System.Net.ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls12;
```



## 参考

- <https://microsoft.com/powershell>
- [PowerShell 中文博客](https://www.pstips.net/)



