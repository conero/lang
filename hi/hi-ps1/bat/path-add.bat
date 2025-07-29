@echo off
::========================================================================
::  Windows 批处理脚本：向 hosts 文件添加域名映射
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
@REM set "DOMAIN1=b6b4cc712a92108d.qaxcloudwaf.com"
@REM set "DOMAIN2=www.bigdata-expo.cn"
@REM set "IP=127.0.0.1"

@REM set "DOMAIN1=b6b4cc712a92108d.qaxcloudwaf.com"
set "DOMAIN1=www.bigdata-expo.cn"
set "IP=b6b4cc712a92108d.qaxcloudwaf.com"

:: 检查是否已存在（避免重复添加）
findstr /i /c:"%IP% %DOMAIN1%" "%HOSTS%" >nul 2>&1
if %errorLevel% equ 0 (
    echo [提示] %DOMAIN1% 已存在于 hosts 文件中，跳过。
) else (
    echo "# hosts set %IP% %DOMAIN1%" >> "%HOSTS%"
    echo %IP% %DOMAIN1% >> "%HOSTS%"
    echo [成功] 已添加: %IP% %DOMAIN1%
)

@REM findstr /i /c:"%IP% %DOMAIN2%" "%HOSTS%" >nul 2>&1
@REM if %errorLevel% equ 0 (
@REM     echo [提示] %DOMAIN2% 已存在于 hosts 文件中，跳过。
@REM ) else (
@REM     echo %IP% %DOMAIN2% >> "%HOSTS%"
@REM     echo [成功] 已添加: %IP% %DOMAIN2%
@REM )

echo.
echo 所有操作已完成！
echo.
pause
