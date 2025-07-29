@echo off
::========================================================================
::  Windows ������ű����� hosts �ļ��Ƴ�ָ������
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
set "TEMP_HOSTS=%temp%\hosts.tmp"
set "DOMAIN1=b6b4cc712a92108d.qaxcloudwaf.com"
set "DOMAIN2=www.bigdata-expo.cn"

echo ���ڴ� hosts �ļ����Ƴ�ָ������...


:: ʹ�� findstr /v ���˵������������������У������ִ�Сд��
findstr /v /i "%DOMAIN1% %DOMAIN2%" "%HOSTS%" > "%TEMP_HOSTS%"

:: ����Ƿ���ƥ����б��Ƴ�
set removed=0
findstr /i "%DOMAIN1%" "%HOSTS%" >nul && set removed=1
findstr /i "%DOMAIN2%" "%HOSTS%" >nul && set removed=1

:: �滻ԭ�ļ�
copy /y "%TEMP_HOSTS%" "%HOSTS%" >nul 2>&1

:: ������ʱ�ļ�
del "%TEMP_HOSTS%" >nul 2>&1

if %removed% equ 1 (
    echo [�ɹ�] ���Ƴ� b6b4cc712a92108d.qaxcloudwaf.com ��/�� www.bigdata-expo.cn ��ӳ�䡣
) else (
    echo [��ʾ] δ�ҵ���ؼ�¼�������Ƴ���
)

echo.
echo ������ɣ�
echo.
pause
