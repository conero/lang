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

### 简介
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



### 注册表

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
- HKEY_LOCAL_MACHINE 
  - *地计算机硬件数据*






> \HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Explorer\FileExts

更改文件后缀名的默认软件打开方式



### 应用自动启动

> 查看方式 1

- *`设置>应用>启动`*  可查看/设置启动的应用





### windows 快捷键

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



### 文件结构

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



### 常用命令

#### ls 文件/目录列表(list)

#### clear	清除命令行/输出台

#### 目录管理

1. `cd` 			目录
2. `mkdir`                   新增目录
3. `rmdir`                   删除目录
4. `mv`                         目录移动
5. `cp`                         文件/目录复制
6. `dirs`                     显示目录
7. `pwd`                       以绝对路径的方式显示用户当前工作目录
8. `cat <file>`                       读取整个文件(*tac* 反序读取文件)
   1. `cat --help  查看命名帮助`
9. `less <file>`      读取部分文件内容
10. `head <file>`       读取头部部分行到屏幕中
11. `tail <file>`       与*head*反之，尾部

#### 进程管理

> `ps` (process)

*ps [options]*

- `$ ps-eo pid,comm,cmd  `  *(-e表示列出全部进程，-o pid,comm,cmd表示我们需要PID，COMMAND，CMD信息)*



#### sudo

> `sudo apt-get install`



#### 应用管理

##### unbantu

>`dpkg ` 查看安装的软件



### vim

> 常用命令

- `:help`			帮助
- `:q`			        退出(quit)
  - `:q!`		取消并退出，不保存
  - `:qw`                退出且保存





# 附录

## 参考

1. [Windows 中的键盘快捷方式](https://support.microsoft.com/zh-cn/help/12445/windows-keyboard-shortcuts)