# Version = PowerShell 7.4.0
# 请求使用管理权限
if (!([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] 'Administrator')) {
    # Start-Process -FilePath PowerShell.exe -Verb Runas -ArgumentList "-File `"$($MyInvocation.MyCommand.Path)`"  `"$($MyInvocation.MyCommand.UnboundArguments)`""
    Start-Process -FilePath pwsh.exe -Verb Runas -ArgumentList "-File `"$($MyInvocation.MyCommand.Path)`"  `"$($MyInvocation.MyCommand.UnboundArguments)`""
    Exit
}

# 等待秒
function wait-sec() {
    Write-Output "";
    Write-Host -ForegroundColor Gray "8s 后退出";
    Start-Sleep 8
}

function check-java-path($path) {
    if (-not (Test-Path $path)) {
        Write-Output ""
        Write-Host -ForegroundColor Red " 糟糕，您的java17安装目录不存在，请检查并重新设置"
        $path = Read-Host -Prompt "请您输入已安装java.exe所在的目录？"
        # $path = Convert-Path $path.Trim()
        $path = $path.Trim()
        $path = Convert-Path $path
    }

    # 检查
    if (-not (Test-Path $path)) {
        $path = check-java-path $path;
    }
    return $path;
}

#
# java 版本切换
# 2024年1月29日 星期一
# Joshua Conero

# 输出java版本号
java -version

# 读取/本届安装的java
$fdJava = (where.exe java);
$javaLst = @();
if ($fdJava.GetType().Name -eq "string") {
    $javaLst += $fdJava;
}
else {
    $javaLst = $fdJava;
}


Write-Output ""
Write-Output " 当前设置的java path："
Write-Output $javaLst
for ($i = 0; $i -lt $javaLst.count; $i++) {
    $el = $javaLst[$i];
    $javaLst[$i] = $el.Replace("\java.exe", "");
}

# java8 路径
$java = 'D:\Program Files\java\jdk-17.0.5\bin'
$java = "$(check-java-path $java)";

# 本身无需改变
if (($javaLst.Length -eq 1) -and ($java -in $javaLst)) {
    Write-Output ""
    Write-Host -ForegroundColor Green " 您当前就是 Java环境，无需更变……"
    Write-Host -ForegroundColor DarkCyan " 当前Java17 PATH: $java"
    Write-Host -ForegroundColor DarkCyan (" 当前Java17 JAVA_HOME: " + ($env:JAVA_HOME))
    Write-Output " Bye, Hacker."
    wait-sec
    exit
}

# Get-Item/可能存在异常
# $javaHome = (Get-Item $java).Parent.FullName
$javaHome = Split-Path -Parent $java
Write-Output ""
Write-Host -ForegroundColor DarkCyan (" 将设置Java17 PATH: " + $java)
Write-Host -ForegroundColor DarkCyan (" 将设置Java17 JAVA_HOME: " + $javaHome)
Write-Output " "

# Write-Output "您是否要切换版本为 java8"
$yes = Read-Host -Prompt "您是否确认切换到 Java17 环境？(n/y)"
if (-not($yes -eq "y")) {
    Write-Output ""
    Write-Output " Bye, Hacker."
    wait-sec
    exit
}


Write-Output ""
Write-Output " 正在切换……"


# 内容处理
# $pathNew = @(); 
# 新增切片
$pathNew = [System.Collections.ArrayList]::new(); 

# 变量数组
foreach ($v in "$env:path".Split(';')) {
    $el = $v.Trim();
    if ($el.Length -lt 1) {
        continue;
    }

    if ($el -in $javaLst) {
        continue;
    }

    # 比较/太慢
    # $isContinue = false;
    # foreach($rf in $javaLst){
    #     $rf = $rf.Replace("/", "\");
    #     $tg = $el.Replace("/", "\");
    #     if($tg -in $rf){
    #         $isContinue = true;
    #         break;
    #     }
    # }
    # if($isContinue){
    #     continue;
    # }

    # $pathNew.Add($el);
    $pathNew += $el;
}

# 追加java8
# $pathNew.Add($java);
$pathNew += $java + ";";



Write-Output "";
$pathStr = ($pathNew -join ";");
Write-Host -ForegroundColor DarkGray $pathStr;
# 设置当前环境变量
$env:Path = $pathStr;

# 环境变量设置
[Environment]::SetEnvironmentVariable('PATH', $pathStr, 'Machine');
[Environment]::SetEnvironmentVariable('JAVA_HOME', $javaHome, 'Machine');


Write-Output "";
java -version
Write-Output "";
Write-Host -ForegroundColor DarkGreen  "设置成功：";
Write-Host -ForegroundColor DarkGray $env:Path;

wait-sec