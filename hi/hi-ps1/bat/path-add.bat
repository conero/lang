@echo off
::========================================================================
::  Windows ������ű����� hosts �ļ��������ӳ��
::  ��Ҫ����ԱȨ������
::========================================================================

:: ����Ƿ��Թ���ԱȨ������
net session >nul 2>&1
if %errorLevel% neq 0 (
    echo.
    echo ���󣺴˽ű������Թ���ԱȨ�����У�
    echo.
    echo ���Ҽ��������ļ���Ȼ��ѡ���Թ���Ա������С���
    echo.
    pause
    exit /b
)

:: �������
set "HOSTS=%SystemRoot%\System32\drivers\etc\hosts"
@REM set "DOMAIN1=b6b4cc712a92108d.qaxcloudwaf.com"
@REM set "DOMAIN2=www.bigdata-expo.cn"
@REM set "IP=127.0.0.1"

@REM set "DOMAIN1=b6b4cc712a92108d.qaxcloudwaf.com"
set "DOMAIN1=www.bigdata-expo.cn"
set "IP=b6b4cc712a92108d.qaxcloudwaf.com"

:: ����Ƿ��Ѵ��ڣ������ظ���ӣ�
findstr /i /c:"%IP% %DOMAIN1%" "%HOSTS%" >nul 2>&1
if %errorLevel% equ 0 (
    echo [��ʾ] %DOMAIN1% �Ѵ����� hosts �ļ��У�������
) else (
    echo "# hosts set %IP% %DOMAIN1%" >> "%HOSTS%"
    echo %IP% %DOMAIN1% >> "%HOSTS%"
    echo [�ɹ�] �����: %IP% %DOMAIN1%
)

@REM findstr /i /c:"%IP% %DOMAIN2%" "%HOSTS%" >nul 2>&1
@REM if %errorLevel% equ 0 (
@REM     echo [��ʾ] %DOMAIN2% �Ѵ����� hosts �ļ��У�������
@REM ) else (
@REM     echo %IP% %DOMAIN2% >> "%HOSTS%"
@REM     echo [�ɹ�] �����: %IP% %DOMAIN2%
@REM )

echo.
echo ���в�������ɣ�
echo.
pause
