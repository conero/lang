@REM 使用跳转程序，过滤一下的注释内容
goto JoshuaConeroStart

@REM Joshua Conero. (20210205)
@REM Window bat simple example. `@` 与 `echo off` 类似可以不使其输出东西。

@REM 此处注释，url 会提示系统错误。 
::         参考： 
@REM https://www.cnblogs.com/zhaoqingqing/p/4620402.html Windows 批处理(bat)语法大全
@REM https://blog.csdn.net/wh_19910525/article/details/8125762    bat批处理的注释语句
@REM https://www.jb51.net/article/21568.htm BAT 特殊符号总结
@REM https://blog.csdn.net/baokx/article/details/14000265 【技巧】批处理,Bat中获取用户输入的指定内容

@REM 接收外部变量（第一个参数）。 `%n` 获取地 n 个参数。
@REM 没有 elseif




:JoshuaConeroStart
@ECHO OFF
@REM @ECHO OFF    关闭命令行输出。

@REM 标题
echo ...Joshua Conero. Write on Date 20210105...

set number=0 
set action="%1"
goto start

:start
if %action% == "" GOTO noParam1

if %action% == help (
    echo "help     输入帮助信息."
    echo "exit     退出系统."
    echo "dir      目录显示."
) else (
    if %action% == ^exit goto ^exit
    if %action% == ^dir goto action_dir
    
    echo "   %action% 命令不存在呢！"
)

goto move_start


:noParam1
echo 没有命令行参数%number%
echo ...
set /a number=%number%+1
goto move_start

:move_start
set /p action=输入命令：
goto start

:exit
@REM 退出处理

@REM 暂定退出
Pause

exit



@REM action 代码
:action_dir
set vcd=%cd%
echo " 当前目录为："%vcd%
echo 当前盘符：%~d0
echo 当前盘符和路径：%~dp0
echo 当前盘符和路径的短文件名格式：%~sdp0
echo 当前批处理全路径：%~f0
echo 当前CMD默认目录：%cd%

goto move_start