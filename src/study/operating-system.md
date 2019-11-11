# 操作系统(OS/Operating System)
>  基本信息
    作者： Joshua Conero
    日期： 2017年6月20日 星期二



## 目录

1. [简介](#menu_overview)
2. [UNIX 1970](#menu_unix)
2. [Windows 1985](#menu_windows)
2. [Linux 1991](#menu_linux)



# <span id="menu_overview">简介</span>



# <span id="menu_unix">UNIX</span> 1970



## 简介



> 特征
>
> > 多任务、多用户

> 日期 / 作者
> > 1969 / 美国AT&T公司的贝尔实验室
> >
> > > 主要作者：Ken Thompson,Dennis Ritchie,Douglas McIlroy

- ![UNIX进化史](./Unix_history-simple.svg)
- ![UNIX进化时间表](./Unix-history.svg)
- 只有符合单一UNIX规范的UNIX系统才能使用UNIX这个名称，否则只能称为类UNIX（UNIX-like）
- UNIX-like 类UNIX
- “一切皆是文件”是 Unix/Linux 的基本哲学之一
- 以网络为核心的设计思想
- 1970-01-01 00:00:00  (unix 元年-1970)



# <span id="menu_windows">Windows</span> 1985
- 创始人: Bill Gates
- Microsoft-DOS -> windows
- 特征
    - 图形化模式 GUI
    - 内核： Windows NT kernel (ntoskrnl.exe)





> 个人使用过的操作系统版本

    Windows XP
    Windows 7.0
    Windows 8.0
    Windows 10.0



> *项目文档*

- *[官方文档](https://docs.microsoft.com/zh-cn/windows/)*
- [参考文档](https://msdn.microsoft.com/zh-cn/library/)



## 注册表 

>  _**Registry**_

- `regedit`       *注册表（register edit）*
- `reg`               *windows 自带的注册表编辑器*



*是Microsoft Windows中的一个重要的数据库，用于存储系统和应用程序的设置信息*

*HKEY_CLASSES_ROOT和HKEY_CURRENT_CONFIG中存放的信息都是HKEY_LOCAL_MACHINE中存放的信息的一部分，而HKEY_CURRENT_USER中存放的信息只是HKEY_USERS存放的信息的一部分*



> 集合包含关系

- HKEY_LOCAL_MACHINE
  - HKEY_CLASSES_ROOT
  - HKEY_CURRENT_CONFIG
- HKEY_USERS
  - HKEY_CURRENT_USER



> **根键的作用** 

- HKEY_USERS
  - *保存了存放在本地计算机口令列表中的用户标识和密码列表,每个用户的预配置信息都存储在HKEY_USERS根键中。HKEY_USERS是远程计算机中访问的根键之一*
- HKEY_CURRENT_USER 
  - *当前登录的用户信息*
- HKEY_CURRENT_CONFIG 
  - *前用户桌面配置(如显示器等)的数据,最后使用的文档列表（MRU）和其他有关当前用户的Windows 98中文版的安装的信息*
- HKEY_CLASSES_ROOT
  - *其文件类型的名称*
  - *存储Windows可识别的文件类型的详细列表，以及相关联的程序。*
- HKEY_LOCAL_MACHINE 
  - *本地计算机硬件数据*




> \HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Explorer\FileExts

更改文件后缀名的默认软件打开方式



> *注册表类型分类*

- `REG_SZ`    *字符串文本*
- *`REG_BINARY`*   *二进制文本*
- `REG_MULTI_SZ`    *多字符串值：含有多个文本值的字符串*
- `REG_DWORD`     *双字值；一个32位的二进制值，显示为8位的十六进制值。*



*URL Protocol： 使用`URL`地址打开应用如：`tencent://`。如在web浏览器中使用 URL 协议打开应用: ` window.open("tencent://message/?uin=346915968")`*



## 应用自动启动

> 查看方式 1

- *`设置>应用>启动`*  可查看/设置启动的应用





## windows 快捷键

> win+

| 快捷键            | 助记                      | 说明                         |
| ----------------- | ------------------------- | ---------------------------- |
| win               | windows                   | 打开或关闭“开始”屏幕         |
| win+A             | windows Action            | 打开操作中心                 |
| win+B             | windows Blur Notification | 将焦点设置到通知区域         |
| win+D             | windows Display           | 显示和隐藏桌面               |
| win+Alt+D         | windows display date      | 显示和隐藏桌面上的日期和时间 |
| win+E             | windows Explorer          | 打开文件资源管理器           |
| win+I             | windows Items             | 打开“设置”                   |
| win+K             |                           | 打开“连接”快速操作(蓝牙)     |
| win+L             | windows Lock              | 锁定你的电脑或切换帐户       |
| win+M/win+Shift+M |                           | 最小化所有窗口(和 win+D 像)  |
| win+P             | windows Play              | 选择演示显示模式             |
| win+R             | windows Run               | 打开“运行”对话框             |
| win+S             | windows Search            | 打开搜索                     |
| win+T             | windows Tabbar            | 循环浏览任务栏上的应用       |
| win+X             |                           | 打开“快速链接”菜单           |
| win+,(逗号)       |                           | 临时速览桌面                 |
| win+Tab           |                           | 打开任务视图                 |



> 其他

| 快捷键            | 助记                      | 说明                         |
| ---- | ---- | ---- |
| *__复制、粘贴及其他常规的键盘快捷方式__* |      |      |
| Ctrl+C/ Ctrl+Insert | Copy | 复制选定文本 |
| Ctrl+V/ Shift+Insert | Paste | 粘贴选定文本 |
| Ctrl+X | Cut | 剪切选定项 |
| Ctrl+Z | Undo | 撤消操作 |
| Alt + Tab | | 在打开的应用之间切换 |
| Ctrl + A | All | 选择文档或窗口中的所有项目 |
| Ctrl + D（或删除键） | Delete | 删除所选的项目，将其移至回收站 |
| Ctrl + R（或 F5） | Refresh | 刷新活动窗口 |
| Ctrl + Y | Redo | 恢复操作 |
| Ctrl + Shift（及箭头键） | | 选择文本块 |
|  | |  |
| *__文件资源管理器键盘快捷方式__* | |  |
| Alt+D | | 选中地址栏 |
| Ctrl+F/Ctrl+E | Find | 选择搜索框 |
| Ctrl+N | New | 打开新窗口 |
| Ctrl+W | | 关闭活动窗口 |
| Ctrl+Shift+N | | 创建新文件夹 |
| Alt + P |  | 显示预览面板 |
| Alt + Enter | | 打开选定项目的“属性”对话框 |





**文件/目录复制地址方法**

```powershell
# 选择【复制为路劲(A)】
shift+右键
```





## CMD

*Windows 命令工具有 `cmd.exe`、`powershell(ps)` 等，其可通过外壳程序调用系统内核相关的程序。*



## 系统



### 快速路径

- 事件查看器：*用于查看系统级别的日志*







# <span id="menu_linux">Linux</span> 1991

- 作者：Linus Benedict Torvalds
- 发布日期： 1991-10-05
- 特征
    - UNIX-like
    - 基于POSIX和UNIX的多用户、多任务、支持多线程和多CPU的操作系统
    - Linux继承了Unix以网络为核心的设计思想，是一个性能稳定的多用户网络操作系统
    - Linux 操作系统的诞生、发展和成长过程始终依赖着五个重要支柱：
        - UNIX 操作系统、
        - MINIX 操作系统、
        - GNU计划、
        - POSIX 标准
        - Internet 网络
    - Linux的基本思想: 
        - 一切都是文件
        - 每个软件都有确定的用途
    - Linux kernel: Freeminix-like kernel sources for 386-AT



> 发布： 1991.10.05

> 类 *Unix* 系统

- *基于POSIX和UNIX的多用户、多任务、支持多线程和多CPU的操作系统*
- *继承了Unix以网络为核心的设计思想，是一个性能稳定的多用户网络操作系统*
  - *一切都是文件*
  - *每个软件都有确定的用途*
- 重要支柱
  - *Unix 操作系统*
  - *Minix* 操作系统
  - *GUN* 计划
  - *POSIX* 标准
  - *Internet 网络*



## 目录结构

**/：**根目录，所有的目录、文件、设备都在/之下，/就是Linux文件系统的组织者，也是最上级的领导者。

**/bin：**bin 就是二进制（[binary](https://baike.baidu.com/item/binary)）英文缩写。在一般的系统当中，都可以在这个目录下找到linux常用的命令。系统所需要的那些命令位于此目录。

**/boot：**Linux的内核及引导系统程序所需要的文件目录，比如 vmlinuz initrd.img 文件都位于这个目录中。在一般情况下，[GRUB](https://baike.baidu.com/item/GRUB)或[LILO](https://baike.baidu.com/item/LILO)系统引导管理器也位于这个目录。

**/cdrom：**这个目录在刚刚安装系统的时候是空的。可以将光驱文件系统挂在这个目录下。例如：mount /dev/cdrom /cdrom

**/dev：**dev 是设备（[device](https://baike.baidu.com/item/device))的英文缩写。这个目录对所有的用户都十分重要。因为在这个目录中包含了所有linux系统中使用的外部设备。但是这里并不是放的外部设备的驱动程序。这一点和常用的windows,dos操作系统不一样。它实际上是一个访问这些外部设备的端口。可以非常方便地去访问这些外部设备，和访问一个文件，一个目录没有任何区别。

**/etc：**etc这个目录是linux系统中最重要的目录之一。在这个目录下存放了系统管理时要用到的各种配置文件和子目录。要用到的网络配置文件，文件系统，x系统配置文件，设备配置信息，设置用户信息等都在这个目录下。

**/home：**如果建立一个用户，用户名是"xx",那么在/home目录下就有一个对应的/home/xx路径，用来存放用户的主目录。

**/lib：**lib是库（[library](https://baike.baidu.com/item/library)）英文缩写。这个目录是用来存放系统动态连接共享库的。几乎所有的应用程序都会用到这个目录下的共享库。因此，千万不要轻易对这个目录进行什么操作，一旦发生问题，系统就不能工作了。

**/lost+found：**在ext2或ext3文件系统中，当系统意外崩溃或机器意外关机，而产生一些文件碎片放在这里。当系统启动的过程中fsck工具会检查这里，并修复已经损坏的文件系统。有时系统发生问题，有很多的文件被移到这个目录中，可能会用手工的方式来修复，或移到文件到原来的位置上。

**/mnt：**这个目录一般是用于存放挂载储存设备的挂载目录的，比如有[cdrom](https://baike.baidu.com/item/cdrom)等目录。可以参看/etc/fstab的定义。

**/media：**有些linux的发行版使用这个目录来挂载那些[usb](https://baike.baidu.com/item/usb)接口的移动硬盘（包括U盘）、CD/DVD[驱动器](https://baike.baidu.com/item/%E9%A9%B1%E5%8A%A8%E5%99%A8)等等。

**/opt：**这里主要存放那些可选的程序。

**/proc：**可以在这个目录下获取系统信息。这些信息是在内存中，由系统自己产生的。

**/root：**Linux超级权限用户[root](https://baike.baidu.com/item/root)的家目录。

**/sbin：**这个目录是用来存放系统管理员的系统管理程序。大多是涉及系统管理的命令的存放，是超级权限用户root的可执行命令存放地，普通用户无权限执行这个目录下的命令，这个目录和/usr/sbin; /usr/X11R6/sbin或/usr/local/sbin目录是相似的，凡是目录sbin中包含的都是root权限才能执行的。

**/selinux** ：对[SElinux](https://baike.baidu.com/item/SElinux)的一些配置文件目录，SElinux可以让linux更加安全。

**/srv** 服务启动后，所需访问的数据目录，举个例子来说，www服务启动读取的网页数据就可以放在/srv/www中

**/tmp：**临时文件目录，用来存放不同程序执行时产生的临时文件。有时用户运行程序的时候，会产生临时文件。/tmp就用来存放临时文件的。/var/tmp目录和这个目录相似。

**/usr**

这是[linux系统](https://baike.baidu.com/item/linux%E7%B3%BB%E7%BB%9F)中占用硬盘空间最大的目录。用户的很多应用程序和文件都存放在这个目录下。在这个目录下，可以找到那些不适合放在/bin或/etc目录下的额外的工具

**/usr/local：**这里主要存放那些手动安装的软件，即不是通过“新立得”或apt-get安装的软件。它和/usr目录具有相类似的目录结构。让软件包管理器来管理/usr目录，而把自定义的脚本（scripts)放到/usr/local目录下面、。

**/usr/share ：**系统共用的东西存放地，比如 /usr/share/fonts 是字体目录，/usr/share/doc和/usr/share/man帮助文件。

**/var：**这个目录的内容是经常变动的，看名字就知道，可以理解为vary的缩写，/var下有/var/log 这是用来存放系统日志的目录。/var/ www目录是定义[Apache](https://baike.baidu.com/item/Apache)服务器站点存放目录；/var/lib 用来存放一些库文件，比如MySQL的，以及[MySQL](https://baike.baidu.com/item/MySQL)数据库的的存放地。

- */*    更目录
  - *bin*    二进制命令，系统中所有命令都在此可见
  - *boot*   系统内核以及引导系统安装程序
  - *dev*     外设目录
  - *etc*       系统配置文件
  - *home*   用户主目录
  - *lib*        程序依赖的库目录，如 so
  - *mnt*      存放挂载储存设备的挂载目录
  - *proc*     系统的版本信息
  - *var*       变量信息目录
  - *tmp*      临时目录
  - *media*  设备挂载信息
  - *srv*        服务相关目录，无服务器程序包等
  - *usr*       用户部分安装的程序
    - *local*    用户本地相关程序
    - *share*   系统公用资源





## 常用命令

```shell
# 操作系统相关的版本信息
cat /proc/version

# 查看本机的ip地址
curl ifconfig.me

# && 用户两个命令的链接
mkdir jc-hellowpy && cd jc-hellowpy

# 环境变量
env

# 输出 path 变量
echo $PATH


# 权限设置
# 可用于如，屏幕提示数据显示不足够的时的操作
chmod 777 <file_path>
```





### ls 文件/目录列表(list)

- `ll`  命令同 `ls -l`



```shell
# 下者等效
ll
ls -l
```







### clear	清除命令行/输出台

### 目录管理

1. `cd` 			目录

2. `mkdir`                   新增目录

3. `rmdir`                   删除目录

4. `mv`                         目录移动，实现重命名

5. `cp`                         文件/目录复制

6. `dirs`                     显示目录

7. `pwd`                       以绝对路径的方式显示用户当前工作目录

8. `cat <file>`                       读取整个文件(*tac* 反序读取文件)
   
   1. `cat --help  查看命名帮助`
   
9. `less <file>`      读取部分文件内容

10. `head <file>`       读取头部部分行到屏幕中

11. `tail <file>`       与*head*反之，尾部

12. `rm`                          删除文件/目录(可删除非空目录)

    ```shell
    # 删除非空文件夹
    rm -rf <dir>
    rm -r <dir>
    
    # 强制删除文件
    rm -f <file>
    ```

13. `find`                      文件或目录查找工具

14. `df`                          *查看磁盘信息*

15. `du`                          *查看当前目录的磁盘信息*

16. `rename`                  *项目重命名*



```shell
# 查看磁盘信息的命令

df -hl 查看磁盘剩余空间
df -h 查看每个根路径的分区大小
du -sh [目录名] 返回该目录的大小
du -sm [文件夹] 返回该文件夹总M数
du -h [目录名] 查看指定文件夹下的所有文件大小（包含子文件夹）

# 文件、目录查找
find -name <name>

# 目录重命名
# <source_dir_name> => <target_dir_name>
rename <source_dir_name> <target_dir_name> <source_dir_name>

# 查看文件信息
stat <file_path>
```





### 进程管理

```sh
# nohup(no-hangup) 后台执行，并挂起
# 通过php，开启内部服务器
# & 表示后台挂起运行
nohup /usr/bin/php -S 0.0.0.0:9200 &
```



> `ps` (process)

*ps [options]*

- `$ ps-eo pid,comm,cmd  `  *(-e表示列出全部进程，-o pid,comm,cmd表示我们需要PID，COMMAND，CMD信息)*



#### Systemd

参考：[systemd](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)

**systemctrl**

*systemctl是 Systemd 的主命令，用于管理系统。*

```shell
# 重启系统
$ sudo systemctl reboot

# 关闭系统，切断电源
$ sudo systemctl poweroff

# CPU停止工作
$ sudo systemctl halt

# 暂停系统
$ sudo systemctl suspend

# 让系统进入冬眠状态
$ sudo systemctl hibernate

# 让系统进入交互式休眠状态
$ sudo systemctl hybrid-sleep

# 启动进入救援状态（单用户状态）
$ sudo systemctl rescue

# 列出所有可用的Unit
$ systemctl list-unit-files
# 列出所有正在运行的Unit
$ systemctl list-units
# 列出所有失败单元
$ systemctl --failed

# 查看 systemd 的日志
$ journalctl -xe
```





### 压缩建档

```shell
$ tar
# 主要类别
-c Create  -r Add/Replace  -t List  -u Update  -x Extract
# 解压
# 创建压缩包
tar.exe -c [options] [<file> | <dir> | @<archive> | -C <dir> ]
# 列出详情
tar.exe -t [options] [<patterns>]
# 解压
tar.exe -x [options] [<patterns>]


tar -xvf <tar_file> -C <tgt_dir>

# zip 压缩
$ zip
# zip 解压
$ unzip
```





### sudo

> `sudo apt-get install`



> **chmod**  改变目录读写权限 

*777*   为最高权限

```shell
sudo chmod -R 777 <file/direction>
```





> **wget**    web http 下载文件 

```shell
# 下载程序包
wget http://nginx.org/download/nginx-1.16.0.tar.gz

# 安装 make 工具
sudo apt install make
# 安装 gcc 编译器
sudo apt install gcc

```



> **whereis** 查看应用所安装的路径/信息

```shell
whereis gcc
```



> root 用户

```shell
# 设置用户账号
sudo passwd root

# 用户登录
su root
```



### 硬件

```bash
# 查看当前系统的架构
arch

# 查看系统信息
uname --all

```





### 应用介绍

#### PuTTY

*Windows 下远程连接 Linux 系统，命令行工具。*

*SSH、telnet 客户端，开源项目。*



*开发创始人： Simon Tatham。*



> *终端快捷键*

```shell
Ctrl + l(L)   	清屏快捷键，同于 clear 命令
ctrl + w —往回删除一个单词，光标放在最末尾
ctrl + k —往前删除到末尾，光标放在最前面（可以使用ctrl+a）
ctrl + u 删除光标以前的字符
ctrl + k 删除光标以后的字符
ctrl + a 移动光标至的字符头
ctrl + e 移动光标至的字符尾 

reset			完全刷新终端屏幕

# 粘贴/复制
点击右键		  粘贴文本
左键选中		  复制文本

# 命令模糊查看
# 键入命令前缀，两个连续的tab键可查询相关的命令
```



> Windows 文件传输到 Linux 系统

*使用自带的命令工具： `pscp`*

```shell
# pscp 文件名称 Linux用户@IP:目标地址
# pscp <filename> linux-user@ip:targetDir
pscp php-7.3.5.tar.gz root@ip:targetDir

# -v 显示进展信息
# 认证错误时，需要执行 ppk 秘钥文件
# <filename> 位置在 -i 选项之后
pscp -i <ppk_file> <filename> user@host:targetDir
# 实例
pscp -i ~\16.13-Linux.ppk ~\python3-3.6.8-10.el7.x86_64.rpm root@y.x.16.13:/usr/local/cro1911 -v
```





#### WinSCP

_*WinSCP*是一个Windows环境下使用SSH的开源图形化SFTP客户端。同时支持SCP协议。它的主要功能就是在本地与远程计算机间安全的复制文件。_

**Linux <=> Windows 之间的文件传输工具(图形化工具)**



软件提供 `PuTTY` 工具集成，通过`ctrl+P`快捷键快速打开；可通过软件保存密码，并且将密码传递给putty，但是软件默认是不开启的。相关配置如下(*版本 5.15.3*)：

- **保存密码** *【新建会话】-> 【保存】中勾选保存密码*
- **传递密码给putty**  *【选项】 -> 【选项】-> 【集成】-> 【应用程序】-> 【记住会话密码并传递给PuTTY(SSH)(R)】*



## 操作系统

### Ubuntu

>`dpkg ` 查看安装的软件



### CentOS

*Community Enterprise Operating System，中文意思是社区企业操作系统*



> 软件 安装/卸载

```shell
# 更新 yum
$ yum update

# yum 软件安装与卸载
$ yum instal <name>
$ yum remove <name>
```



> 用户目录

```shell
$ /username/
```





> 版本信息查看

```shell
# 查看系统信息
$ uname -a

# 查看版本信息
$ cat /etc/redhat-release
```





> LNMP

```shell
# 安装编译环境
$ yum -y install make zlib zlib-devel gcc-c++ libtool  openssl openssl-devel

# ----------------------【nginx】------------------------
# 安装 nginx
$ wget <loaddown.net.url>
$ tar zxvf nginx-1.6.2.tar.gz	解压
$ cd nginx-1.6.2				目录转移
$ ./configure					编译安装
$ make							安装/ make install

# 启动 nginx
$ /usr/local/nginx/sbin/nginx
# 停止服务器
$ /usr/local/nginx/sbin/nginx -s stop 或 /usr/local/nginx/sbin/nginx -s quick
```



> *其他软件安装*

```shell
# yum 安装/卸载 git 软件
$ yum install git
$ yum remove git 

# 通过源码安装
# 错误  http.h:6:23: fatal error: curl/curl.h: No such file or directory
$ yum install -y libcurl*
$ yum install -y expat*
$ wget https://github.com/git/git/archive/v2.21.0.tar.gz
$ tar -zxvf v2.21.0.tar.gz
$ cd v2.21.0
$ make && make install
```



> _**用户/用户组**_

- *用户列表文件：/etc/passwd/*
- *用户组列表文件：/etc/group*

```shell
# 查看可以登录系统的用户
$ cat /etc/passwd | grep -v /sbin/nologin | cut -d : -f 1

# 新增分组 ancient
$ groupadd ancient

# 添加用户
$ useradd -g ancient chiyou 		# 将 chiyou 添加到 ancient 分组中
$ useradd chiyou 

# 查看当前用户分组
$ group
# 查看 kuafu 用户的所属分组
$ groups kuafu
# id 用户(查看用户所属组）
$ id root
```



#### rpm

*`Red-Hat Package manager`(same like `RPM Package Manager`)*   RPM软件包管理器，由于其的开放式理念已经使用于许多的Linux发行版。

```shell
# i=> install; v = verbose; h = hash
$ rpm -ivh your-package.rpm

# 查看表信息
$ rpm -ql <package>
# 下载软件
# erase (uninstall) package
$ rpm -e <package>
# 更新软件
# upgrade package(s)
$ rpm -U <package>
```



*[rpm 软件下载地址](http://www.rpmfind.net/linux/RPM/)*





## vim

旧版命令 *`vi`*，*`vim`全部兼容前者。`Insert` 插入文本命令, 推荐使用 `vim`命令。*



> 常用命令

- `:help`			帮助
- `:q`			        退出(quit)
  - `:q!`		取消并退出，不保存
  - `:wq`                退出且保存
  - `ZZ`                  退出且保存
- `:e!`                       丢弃修改并打开原文件





> 方位移动

```
  k
h   l
  j  
```

- `0`   *行首(数字)*
- `$`   *行尾*
- `w`   *下一个单词起始处*
- `nw`  *n个单词以后*
- `e`   *单词末尾，与`w`类似*
- `gg` *文件起始处*
- `G`  *文件结尾处*



> 翻页

- `ctrl + e`    *向下滚动一行*
- `ctrl + y`    *向上滚动一行*
- `Ctrl + f`    上一页
- `Ctrl + b`    下一页
- `ctrl + u`    上半页 (up)
- `ctrl + d`    下半页 (down)



> *搜索与替换*

*在命令模式下，输入 `/` 用于文本查找. 从一般模式进编辑模式，只需按i、I、a、A、o、O、r和R中某个键即可。命令模式：输入：或者/即可进入命令模式。*

*`u`       撤销代码， `Ctrl + R`  恢复撤销*





## issue

### setgid: Operation not permitted

> root 用户切换无效

```shell
su root
```









## shell

> Bourne Shell 变量，`$<NAME>`  内部命令

- `HOME`   *当前用户所在主页*
- `PATH`    *环境变量*





### [bash](http://www.gnu.org/software/bash/manual/)

> 内部变量

- `BASH`    *当前正在执行的 bash 实例*
- `BASHPID`   *当前 bash 的进程 ID*
- `BASH_VERSION`    *bash 版本信息*
- `LANG`         *语言信息*





>*常用 Builtin 内建命令*

- `cd`     *改变当前的工作目录(current directory.)*
- `pwd`   *输出当前的工作目录，`$PWD` 对应常量*
- `exit`  *关闭当前的 shell 终端，同`logout`*
- `times`  *输出时间*
- `echo`    *输出到终端*
- `help`    *用于查看内建命令*
- `source <filename>`    *执行指定的，shell 文件*





# 附录

### RTOS

*Real-time operating system , 实时操作系统。是管理系统硬件和软件资源的系统软件，以方便开发者使用，操作系统管理的资源包括处理器、存储器、外设、甚至包括文件系统等等。*



实时操作系统最大的特色就是其“实时性”。也就是说，如果有任务需要执行，实时操作系统会立即（在较短时间内）执行该任务，保证了任务在指定时间内完成。

实时操作系统根据任务执行的实时性，分为“硬实时”操作系统和“软实时”操作系统，“硬实时”操作系统比“软实时”操作系统响应更快、实时性更高，“硬实时”操作系统大多应用于工业领域。

- “硬实时”操作系统必须使任务在确定的时间内完成。
- “软实时”操作系统能让绝大多数任务在确定时间内完成。



## 键盘

*keyboard*



*作为计算机系统的一种 `输入` 设备.*



*`QWERTY` “打字机之父”——美国人克里斯托夫·拉森·肖尔斯（Christopher Latham Sholes）发明.* 





> 基准建， 指法分工实现盲打

- 左手手指 ：`A` `S` `D` `F`    右手手指：  `J` `K` `L` `;`

*键盘背记规则 (顺口溜):*

```
QWERTY 	    七碗鹅肉汤					YUIOP 已无一，阿婆
ASDFG	    爱死豆腐羹    				HJKL 回家快乐
ZXCVBNM		自行车未帮你买
```





## 差异性

### LF/CRLF

- LF:  *Line Feed (`\n`)* 换行，Unix 句末使用换行
- CRLF:  *Carriage Return Line Feed (`\r\n`)*  回车换行，Windows 下表示句尾使用回车换行两个字符
- CR:  *Carriage Return (`\r`)*  回车换行，Mac 下表示句尾使用回车字符





## 参考

1. [Windows 中的键盘快捷方式](https://support.microsoft.com/zh-cn/help/12445/windows-keyboard-shortcuts)