# PowerShell

> 2018年10月19日 星期五




## 简介

*Windows PowerShell® 是基于任务的命令行管理程序和脚本语言，专为进行系统管理而设计。*

*UNIX 系统一直有着功能强大的脚本（shell），Windows PowerShell 的诞生就是要提供功能相当于 UNIX 系统 BASH 的命令行外壳程序，同时也内建脚本语言以及辅助脚本程序的工具。*

*PowerShell 使用“谓词 - 名词”命名系统。 每个 cmdlet 名称都由一个标准谓词、连字符和特定名词组成。*



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





### 输入输出

> 调用 VB 做图形化输入

```powershell
[void][System.Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
$cn = [Microsoft.VisualBasic.interaction]::inputbox('Enter Your Name', 'Name', 'Jc')
echo ">>:  $cn"
```







## 命令参考

```powershell
# 获取命令列表
Get-Command

# 查看命令已经分配的别名
Get-Command -CommandType Alias

# 常用命名
# 打印当前 ps 的版本信息
$psversiontable
$host
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
```



### 文件/目录

*目录存在检测*

```powershell
$path = './runtime/bundle'
if(-not (Test-Path -Path $path )){
    mkdir $path
}
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


# dir 与 其他 ls 命令差不多
dir
```



### 服务/进程

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



对应 C #*

```c#
System.Net.ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls12;
```



## 参考

- <https://microsoft.com/powershell>
- [PowerShell 中文博客](https://www.pstips.net/)



