@echo off
::========================================================================
::  Windows 批处理脚本：从 hosts 文件移除指定域名
::  需要管理员权限运行
::========================================================================

:: 检查是否以管理员权限运行
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo.
    echo 错误：此脚本必须以管理员权限运行！
    echo.
    echo 请右键单击此文件，然后选择“以管理员身份运行”。
    echo.
    pause
    exit /b
)

:: 定义变量
set "HOSTS=%SystemRoot%\System32\drivers\etc\hosts"
set "TEMP_HOSTS=%temp%\hosts.tmp"
set "DOMAIN1=b6b4cc712a92108d.qaxcloudwaf.com"
set "DOMAIN2=www.bigdata-expo.cn"

echo 正在从 hosts 文件中移除指定域名...


:: 使用 findstr /v 过滤掉包含这两个域名的行（不区分大小写）
findstr /v /i "%DOMAIN1% %DOMAIN2%" "%HOSTS%" > "%TEMP_HOSTS%"

:: 检查是否有匹配的行被移除
set removed=0
findstr /i "%DOMAIN1%" "%HOSTS%" >nul && set removed=1
findstr /i "%DOMAIN2%" "%HOSTS%" >nul && set removed=1

:: 替换原文件
copy /y "%TEMP_HOSTS%" "%HOSTS%" >nul 2>&1

:: 清理临时文件
del "%TEMP_HOSTS%" >nul 2>&1

if %removed% equ 1 (
    echo [成功] 已移除 b6b4cc712a92108d.qaxcloudwaf.com 和/或 www.bigdata-expo.cn 的映射。
) else (
    echo [提示] 未找到相关记录，无需移除。
)

echo.
echo 操作完成！
echo.
pause
