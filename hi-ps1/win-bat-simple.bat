@REM ʹ����ת���򣬹���һ�µ�ע������
goto JoshuaConeroStart

@REM Joshua Conero. (20210205)
@REM Window bat simple example. `@` �� `echo off` ���ƿ��Բ�ʹ�����������

@REM �˴�ע�ͣ�url ����ʾϵͳ���� 
::         �ο��� 
@REM https://www.cnblogs.com/zhaoqingqing/p/4620402.html Windows ������(bat)�﷨��ȫ
@REM https://blog.csdn.net/wh_19910525/article/details/8125762    bat�������ע�����
@REM https://www.jb51.net/article/21568.htm BAT ��������ܽ�
@REM https://blog.csdn.net/baokx/article/details/14000265 �����ɡ�������,Bat�л�ȡ�û������ָ������

@REM �����ⲿ��������һ���������� `%n` ��ȡ�� n ��������
@REM û�� elseif




:JoshuaConeroStart
@ECHO OFF
@REM @ECHO OFF    �ر������������

@REM ����
echo ...Joshua Conero. Write on Date 20210105...

set number=0 
set action="%1"
goto start

:start
if %action% == "" GOTO noParam1

if %action% == help (
    echo "help     ���������Ϣ."
    echo "exit     �˳�ϵͳ."
    echo "dir      Ŀ¼��ʾ."
) else (
    if %action% == ^exit goto ^exit
    if %action% == ^dir goto action_dir
    
    echo "   %action% ��������أ�"
)

goto move_start


:noParam1
echo û�������в���%number%
echo ...
set /a number=%number%+1
goto move_start

:move_start
set /p action=�������
goto start

:exit
@REM �˳�����

@REM �ݶ��˳�
Pause

exit



@REM action ����
:action_dir
set vcd=%cd%
echo " ��ǰĿ¼Ϊ��"%vcd%
echo ��ǰ�̷���%~d0
echo ��ǰ�̷���·����%~dp0
echo ��ǰ�̷���·���Ķ��ļ�����ʽ��%~sdp0
echo ��ǰ������ȫ·����%~f0
echo ��ǰCMDĬ��Ŀ¼��%cd%

goto move_start