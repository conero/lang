# 操作系统(OS/Operating System)
>  基本信息
    作者： Joshua Conero
    日期： 2017年6月20日 星期二



# <span id="menu_linux">Linux</span> 1991

- 作者：Linus Benedict Torvalds
- 发布日期： 1991-10-05
- 特征
    - UNIX-like
    - 基于POSIX和UNIX的多用户、多任务、支持多线程和多CPU的操作系统
    - Linux继承了Unix以网络为核心的设计思想，是一个性能稳定的多用户网络操作系统
    - Linux 操作系统的诞生、发展和成长过程始终依赖着五个重要支柱：
        - UNIX 操作系统、
        - MINIX 操作系统
        - GNU计划
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

**/dev：**dev 是设备（[device](https://baike.baidu.com/item/device))的英文缩写。这个目录对所有的用户都十分重要。因为在这个目录中包含了所有linux系统中使用的外部设备。但是这里并不是放的外部设备的驱动程序。这一点和常用的windows,dos操作系统不一样。它实际上是一个访问这些外部设备的端口。可以非常方便地去访问这些外部设备，和访问一个文件，一个目录没有任何区别。（被内核维护的设备列表）

**/etc：**etc这个目录是linux系统中最重要的目录之一。在这个目录下存放了系统管理时要用到的各种配置文件和子目录。要用到的网络配置文件，文件系统，x系统配置文件，设备配置信息，设置用户信息等都在这个目录下。

**/home：**如果建立一个用户，用户名是"xx",那么在/home目录下就有一个对应的/home/xx路径，用来存放用户的主目录。

**/lib：**lib是库（[library](https://baike.baidu.com/item/library)）英文缩写。这个目录是用来存放系统动态连接共享库的。几乎所有的应用程序都会用到这个目录下的共享库。因此，千万不要轻易对这个目录进行什么操作，一旦发生问题，系统就不能工作了。

**/lost+found：**在ext2或ext3文件系统中，当系统意外崩溃或机器意外关机，而产生一些文件碎片放在这里。当系统启动的过程中fsck工具会检查这里，并修复已经损坏的文件系统。有时系统发生问题，有很多的文件被移到这个目录中，可能会用手工的方式来修复，或移到文件到原来的位置上。

**/mnt：**这个目录一般是用于存放挂载储存设备的挂载目录的，比如有[cdrom](https://baike.baidu.com/item/cdrom)等目录。可以参看/etc/fstab的定义。

**/media：**有些linux的发行版使用这个目录来挂载那些[usb](https://baike.baidu.com/item/usb)接口的移动硬盘（包括U盘）、CD/DVD[驱动器](https://baike.baidu.com/item/%E9%A9%B1%E5%8A%A8%E5%99%A8)等等。

**/opt：**这里主要存放那些可选的程序。用户即程序目录，如oracle等

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



**/usr/bin/**  类似系统环境变量



### 动态链接库

> linux默认读取的动态链接库目录有

- **/lib**
- **/usr/lib**
- **/usr/local/lib**



> 环境变量 **LD_LIBRARY_PATH**

```shell
LD_LIBRARY_PATH
```



> 配置文件 **/etc/ld.so.conf** 及其包含的 **/etc/ld.so.conf.d/*.conf**



## 常用命令

```shell
# 操作系统相关的版本信息
cat /proc/version

# 查看 cpu 信息，如核数等
lscpu

# 操作系统相关信息
uname -a

# 查看本机的ip地址（参考  https://ifconfig.me/）
curl ifconfig.me

# 查看本机账号
ifconfig -a

# && 用户两个命令的链接
mkdir jc-hellowpy && cd jc-hellowpy

# 环境变量
# 环境变量，cat /etc/os-release查看系统定义的环境变量
env
#过滤
env | PATH

# 设置环境变量
export env_key=value
# 如设置
export JC_Name='s * ee , sssss。'
#读取变量
echo $JC_Name
# 或使用 printenv 打印
printenv JC_Name

# 输出 path 变量
echo $PATH

# 权限设置
# 可用于如，屏幕提示数据显示不足够的时的操作
chmod 777 <file_path>
# 权限授权给指定的用户
# 将当前目录所有文件与子目录拥有者设置给nginx、群体使用者设置给 nginxgroup
#
# 父级属于其他用户，子目录属于不同用户时相互不影响
chown -R nginx:nginxgroup *

# 内容/文本查找
# 查看当前目录下所有 html 内容<h3>标签
grep "<h3>Joshua Conero<h3/>" *.html


# 查看命令行类型，hashed/builtin/alias
# type $command
type python3

locate *
```



**ls前缀命令**：

1. lscpu：列出 cpu 相关的信息
1. lsinitrd：列出初始化内存盘（initrd）的内容。Initrd是一个小型的临时根文件系统，用于在Linux启动过程中加载必要的驱动程序和模块，以便挂载真正的根文件系统。
2. lsipc：显示系统中的进程间通信（IPC）资源，包括信号量、消息队列、共享内存段等。
3. lslocks：列出当前系统中的锁，包括文件锁、POSIX消息队列锁等。这对于诊断死锁问题很有帮助。
4. lslogins：显示系统中已知用户的登录信息，包括用户名、终端类型、登录时间和ID等。
5. lsmem：显示系统内存的使用情况，包括物理内存、内核使用的内存区域以及NUMA（非统一内存访问）节点的内存信息。
6. lsmd：这个命令的具体含义可能依赖于特定的系统或软件包，但通常可能与系统管理或监控有关，比如在某些系统中可能与监控磁盘健康状态有关。7. lsns：列出系统中的命名空间（namespaces），Linux命名空间是一种隔离技术，允许在单一主机上运行多个独立的用户空间实例。
8. lsblk：列出所有可用的块设备及其属性，如磁盘、分区等，还可以显示挂载点、UUID、设备大小等信息。
9. lsmcli：这个命令可能是特定于某个硬件或管理工具的命令，例如与某些系统的硬件管理有关，但具体功能需要查看对应的文档或软件说明。
10. lsgpio：列出系统中的GPIO（通用输入输出）引脚状态和配置，这对于硬件调试和控制至关重要。
11. lsof：列出当前系统中打开的文件，包括进程ID、用户、文件描述符等信息。这个命令对于找出哪些进程正在占用文件或端口非常有用。
12. lsiio：这个命令可能是用来查询I/O（输入/输出）设备信息的，但具体功能依据上下文可能有所不同，可能涉及硬件管理或系统性能监控。
13. lshw-gui：硬件清单（Hardware Lister）的图形界面版本，用于展示详细的硬件配置信息，包括CPU、主板、内存、网络设备等。
14. lscgroup：列出控制组（cgroups）的信息，cgroups是Linux内核的一个特性，用于限制、记录和隔离进程组使用的资源。
15. lssubsys：这个命令可能用于列出系统中的子系统信息，具体用途可能与内核子系统、设备驱动或特定软件框架相关，但没有一个广泛认可的标准命令与此名称对应，具体功能可能依赖于特定环境或自定义脚本。
16. lsscsi：列出系统中的SCSI设备，包括磁盘、光驱、磁带机等，提供设备的SCSI ID、类型、LUN（逻辑单元号）等信息。
17. lsusb：列出连接到系统的USB设备，包括设备厂商、产品ID、序列号和设备描述等信息，对于识别和管理USB设备很有帮助。





### ls 文件/目录列表(list)

- `ll`  命令同 `ls -l`



```shell
# 下者等效
# 查看文件的读写能力
ll
ls -l

# 显示所有文件列表，包括隐藏的文件
ls -a

# 用户目录
ls ~

# 根据文件夹大小排名
ls -S

# 安装时间排序
# 时间顺序
ll -rt
# 时间倒序(修改时间)
ll -t
# 多路径
ll /path1/ /path2/ ...
# 人可读格式显示，如文件大小的人性化处理
ll -h
# drwxr-xr-x  8 conero conero 4.0K Mar 20 19:36 .git/
# -rw-r--r--  1 conero conero   84 Mar 19 16:55 .gitignore
# 首字符含义: - 普通文件，d 为目录，l为链接符号（表软链接等）。
```



通配符使用

- `*`    任意匹配规则
- `?`    自匹配一次
- `[]`  类似规则，`[rv]` 支持含“r”或“v”的字符串，`[0-9a-z]` 范围匹配





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
    # -r 表示递归清除
    rm -rf <dir>
    rm -r <dir>
    
    # 强制删除文件
    rm -f <file>
    
    rm -rf *.py runtime many values
    ```

13. `find`                      文件或目录查找工具

14. `df`                          *查看磁盘信息*

15. `du`                          *查看当前目录的磁盘信息*

16. `rename`                  *项目重命名*



cd 改变当前的目录

```shell
# 跳转至用户目录
cd 

# 跳转值先前的目录中
cd -
```



file 命令

```shell
# 显示 composer.json 的文件类型
file composer.json
# composer.json: JSON data
```





其他命令

```shell
# 查看磁盘信息的命令
df -hl          #查看磁盘剩余空间
df -h           #查看每个根路径的分区大小
df -h --total   #查看并统计总的文件大小
df -hT          #显示
du -sh [目录名]  #返回该目录的大小，以及挂载情况

du -sm [文件夹]  #返回该文件夹总M数
du -h [目录名]   #查看指定文件夹下的所有文件大小（包含子文件夹）

# 查看当前顶级目录下文件大小，并排逆序
du -sh /* | sort -hr

# 显示占用空间最大的前10个顶级目录
du -sh /* | sort -hr | head -n 10

# 文件、目录查找
# 文件名查找
# -iname 忽略大小写
# -regex 正则表达式
find -name <name>
# 查看 [/] 目录下的 py 文件
find / -name "*.py"
# 查找当前目录下不以 "py" 为后缀的文件
find . ! -name "*.sh"
# 查看包含【modules】的名称
find / -path "*modules*"

# 查找目录下的“m3u8”文件并删除目录
find D:/conero/phpapps/website/iCloud/uc/uploads/ -name "*.m3u8" | xargs rm -fr

# 查找顶级目录，且仅展示文件类型
find /etc/nginx/conf.d/ -type f  -maxdepth 1
# 查找文件并将其复制到指定目录，如备份
find /etc/nginx/conf.d/ -type f  -maxdepth 1 | xargs cp -t /etc/nginx/conf.d/backup-240117
# 查找目录下，大小为 100kB-10MB 的文件（k/KiB,M/MiB,G/MiG）
# +表大于，-小于，无+/-表等于
# atime   最近访问时间；ctime   最新修改inode的时间；mtime   最新文件修改时间
# amin    类似，单位为分钟
find ./upload/ -size '+100k' -size '-10M'
# 查找修改为在 [1-10]分钟的文件列表
find -mmin +1 -mmin -10

# 目录重命名
# <source_dir_name> => <target_dir_name>
rename <source_dir_name> <target_dir_name> <source_dir_name>
# 使用 mv 亦可来重命名
mv <old-name> <new-name>
# 将当前目录所有文件/目录移动到上级
mv * ..

# 查看文件信息
stat <file_path>

# 文件覆盖
cp -rf ./svn-173/* ./
cp -r ./dir-all/. ./new-dirall
# 复制文件 fl1 --> fl2. fl2 不存在是创建否则覆盖。
cp fl1 fl2
# 复制当前目录下所有的html文件
cp -u *.html ../copt-html
# 重新提醒，输入: y  才行

# 相对地址转绝对地址
readlink -f ./

# 多个目录同时创建
mkdir -p /opt/new/test/01

# 用户实时加载日志尾部行（默认10行，-n 10）
# 可用于日志实时查看
tail -f /home/logs/nginx/error.log

# 查看文件大小
ll /home/logs/nginx/ -h
# 同名
ls -l --color=auto -h
```



计算目录下文件sha值

```shell
# 读取目录下文件 sha
find ./ -type f -print0 | xargs -0 sha256sum

# 读取目录下文件 md5
find ./ -type f -print0 | xargs -0 sha256sum

# 使用 [-maxdepth 1] 设置文件的文件文件几倍
find ./ -type f -print0 -maxdepth 1  | xargs -0 sha256sum
```



磁盘挂载

```shell
# 查看磁盘分区
# 列举所有系统盘设备及其分区信息
fdisk -l

# 查看磁盘被表信息
lsblk -l
# 通过看，文件系统类型是否正常显示，若未显示需进行格式化
lsblk -f

# 如将 /dev/sda1 挂载到 /mnt/newdisk
mkdir /mnt/newdisk
mount /dev/sda1 /mnt/newdisk

# mount 手动挂载，在服务器重启后会回丢失
# 因此需要，通过修改 /etc/fstab 实现重启后自动挂载
# 查看得 uuid 即可,或有 lsblk -f 查看
blkid /dev/sdb
# 修改配置文件，修改配置文件后执行 `mount -a` 即可根据配置挂载或重启后
vi /etc/fstab
#UUID=f53746a1-3b7f-449b-a670-c16b733c75fb /opt/systemlist      ext4    defaults        0 0

# 部分磁盘挂载是提示应用磁盘只读，则需先进行格式化磁盘
mkfs.ext4 /dev/sdb

# 将 iso 文件挂载到指定目录
mount -o loop KingbaseES_V008R006C008B0020_Aarch64_install.iso /opt/Kingbase/install/
# 卸载 iso 文件
umount /opt/Kingbase/install/
```



**lsblk** 管理和查看块设备的信息

```shell
# 管理和查看块设备的信息
# （树形）显示所有块设备的基本信息
lsblk

# 显示完整设备路径
lsblk -p
```



替代命令 [fd](https://github.com/sharkdp/fd) 其也可快速查找文件

```shell
# 默认显示文件列表，所有（自己）
fd

# 查找 c 盘中，大于等于10M的log文件
fd -S '+10m' -l -g "*.log" 'c:\'

# 查找 e 盘中，含snms的名称
fd -i -g "*snms*" 'E:\'

# 使用 d 限制级别
fd -i -g "*snms*" -d 1
```



用户权限查看

```shell
# 查看用户 nobody 是否可访问目录
sudo -u nobody ls -l ./path/dirname

# 设置nobody，用于某一用户组
# 如设置必须保证父级用于此权限，否则无法访问
sudo chown -R nobody:nobody /opt/www/loan-syste
```



路径协助

```shell
# realpath 相对路径转绝对路径
realpath ./redis.cof
```



> ln **软链接**

```shell
# ln 必须使用绝对地址（经实际错误时无效）
# -s 为实际的目标路径（源目录）
ln -s <源目录> <软连接目录>
ln -s /opt/systemlist/yzm/uploads /opt/systemlist/yzm/public/uploads
# 如果资源已存在则强制覆盖
ln -sf /opt/systemlist/yzm/uploads /opt/systemlist/yzm/public/uploads
# 输出详细信息
ln -sv /opt/systemlist/yzm/uploads /opt/systemlist/yzm/public/uploads
```



#### 内容查看或管理

```shell
# 读取文件
# 打开数据文件
cat Readme.md

# 只能向前翻页
# 从前向后分页显示文件内容。
more Readme.md

# less 是对早期more的改进
less Readme.md
```



#### 环境变量

输出shell 中环境变量

```shell
# 输出环境变量，即linux中的应用直接读取环境变量的参数
echo $PATH

#
# 配置全局环境变量
# 1.在 /etc/profile.d/ 下新一个 shell 文件
vi /etc/profile.d/java_home.sh
# 内部编辑
# #!/bin/bash
# export JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64 
# 2. 保存后，执行
source /etc/profile

# 指定用户设置特定环境变量
vi ~/.bash_profile
# 附加环境变量
export PATH=/opt/gypc_ams/php-7.2.34/sapi/cli/:/opt/gypc_ams/php-7.2.34/sapi/fpm/:$PATH
source ~/.bash_profile
```



### 服务(进程)管理

linux  服务管理一般由 **service** 和 **systemd** 两个命令管理，后者出现晚于前置，因此其基本上后向兼容与它。



服务目录

- service 对应的init.d目录： `/etc/init.d/`

- system 脚本目录：`/lib/systemd/system/`



service，`unit.service` 语法，如配置nginx服务

```toml
[Unit]
# 服务基本信息
#

# 服务名称
Description=nginx

#服务启动先决条件
After=networks.target

[Service]
# 服务启动脚本
#

Type=forking

# 服务启动脚本
ExecStart=/usr/local/nginx/sbin/nginx

# 服务停止
ExecStop=/usr/local/nginx/sbin/nginx -s stop

# 服务重启
ExecReload=/usr/local/nginx/sbin/nginx -s reload

ExecStartPre=/usr/local/nginx/sbin/nginx -t

PrivateTmp=true

[Install]
# 服务安装信息
#

#服务启动方式
# multi-user.target 多用户方式启动，服务自动启动服务。
WantedBy=multi-user.target

#RequiredBy=

#Also=

#Alias=
```



相公控制命令：

```shell
# 查看服务列表
systemctl list-units

# 已启动的服务列表
systemctl list-units --type=service

# 启动服务
systemctl start nginx

# 停止服务
systemctl stop nginx

# 服务详情，可查看服务器是否运行自动启动
systemctl status nginx

# 更改某些unit配置文件后，需重新加载配置文件
systemctl daemon-reload

# 服务开启重启（enable）
# 其他，修改前执行  deamon-reload
#    disable 服务未运行
#    masked  服务不可运行
#    disable 服务未运行
#    static  只有在别的单元启动时才被使用
systemctl enable nginx
```







**nohup** 挂起应用

```sh
# nohup(no-hangup) 后台执行，并挂起
# 通过php，开启内部服务器
# & 表示后台挂起运行
nohup /usr/bin/php -S 0.0.0.0:9200 &

# 查找进程，是否运行或者杀死进程
ps -aux|grep <name>
# root      5445  0.0  0.5 115052  9424 ?        Sl   Nov15   0:02 ./wedding
# 进程号如： 5445
# 查看所有的进程
ps -A

kill -9 <进程号>
# 关闭所有的 php-fpm 进程
killall php-fpm

#查询进程炳杀掉它，如进程: yangsu
ps -ef|grep yangsu | grep -v grep | awk '{print $2}' | xargs kill -9

# 根据 pid 获取进程执行的详细情况
ps -p 25647 -f
ps -ef|grep yangsu
```



Linux 服务端端口号占用情况查询，ps 命令可查询启动它的用户信息：

```shell
# 1) 查询进程，并获取 pid
# xxx 为二进制名称
ps -ef|grep xxx
# 2) 使用pid查询的程序
netstat -anp|grep <pid>

# 方法二
# 1) 查看端口号的占用 pid，第二部同上所示
lsof -i:<port>

# 其他
netstat -lnt|grep 9000

# 查看当前服务器监听的tcp/upd服务
netstat -tunlp
netstat -tunpl | grep 9000

# 当前的端口开发情况
# udp 协议
netstat -nupl
# tcp 协议
netstat -ntpl
```



**service**

```shell
#系统所在目录
#/etc/init.d/
# 可直接编写服务脚本。

# 常见的对应服务名称: `service 服务名 [start | stop | restart | reload | status]`
# service命令启动redis脚本
service redis start
# 直接启动redis脚本
/etc/init.d/redis start

# 开机自启动
update-rc.d redis defaults
```







> `ps` (process)

*ps [options]*

- `$ ps-eo pid,comm,cmd  `  *(-e表示列出全部进程，-o pid,comm,cmd表示我们需要PID，COMMAND，CMD信息)*



```shell
# 指定 pid = 2356,并查看其二进制路径等
ps -fp 2356 -l -www

# 进程内存情况使用排名
ps aux --sort=-%mem
# 进行名称查询
ps -aux |grep nginx
```



#### top

top 命令是Linux下常用的性能分析工具，能够实时显示系统中各个进程的资源占用状况，类似于Windows的任务管理器。

```shell
# 查看当前 cpu/磁盘 使用率
top
```



#### Systemd

system daemon，linux下一种init软件。致力于取代 `initd` 

参考：[systemd](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)   

主要命令如：

- systemctl                          systemd 主命令
- systemd-analyze             查看启动耗时
- hostnamectl                    查看当前主机的信息
- localectl                            查看本地化配置信息
- timedatectl                      查看当前时区信息
- loginctl                             查看当前用户登录信息
- journalctl                         日志信息查看



```shell
# 查看服务启动时间
systemd-analyze

# 查看各应用启动时间
systemd-analyze blame

# 查看当前主机信息
hostnamectl

# 查看系统当前本地化配置
localectl

# 查看系统的当前时区信息
timedatectl

# 查看当前用户登录信息
loginctl
```



##### systemctl

*systemctl是 Systemd 的主命令，用于管理系统。*

`/lib/systemd/system/` 为 Unit 单元所在目录。

```shell
# 重启系统
sudo systemctl reboot
# 执行系统重启
reboot
# 系统关机
shutdown 

# 关闭系统，切断电源
sudo systemctl poweroff

# CPU停止工作
sudo systemctl halt

# 暂停系统
sudo systemctl suspend

# 让系统进入冬眠状态
sudo systemctl hibernate

# 让系统进入交互式休眠状态
sudo systemctl hybrid-sleep

# 启动进入救援状态（单用户状态）
sudo systemctl rescue

# 列出所有可用的Unit
systemctl list-unit-files
# 查看已启动的服务
systemctl list-unit-files --state=enabled
# 列出所有正在运行的Unit
systemctl list-units
# 列出所有失败单元
systemctl --failed

# 使用 systemctl 启动服务
systemctl restart nginx

# 打印系统状态
systemctl status

#打印显示的全部进程
systemctl

# 查看 systemd 的日志
journalctl
journalctl -xe
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

#
# 创建 tar.gz 压缩文件
# ./dist 为顶级目录
tar -czvf path-dist.tar.gz ./dist
# ./dist 不做顶级，加上参数`-C` 和 `.`
tar -czvf path-dist.tar.gz -C ./dist .

# zip 压缩
$ zip
# zip 解压
$ unzip
# 如，解压 zip 包到当前目录
unzip -v ./upgrade-list/blockd_ow-zzt-211020-02-failure.zip
# 其他选项：
# 	-d <dir> 压缩到指定的目录
# 	-l 查看压缩的文件
# 	-t 检查文件
#	-o 默认覆盖所有文件（保存安全前提下请备份文件）
# 帮助文: info unzip

# 解压并默认覆盖已经存在的文件
unzip -o file_name.zip
```



### 内存以及存储管理

#### 内存

- free               系统内存使用情况显示

```shell
# 系统的可用内存情况
# 以 “MB” 为单位查询内存使用情况
free -m
# 以 “GB” 为单位显示
free -g
# 以人们可读的方式显示数据
free -h

# 进程使用内存（CPU/磁盘）情况
top
# 进入界面后
# 键入"1"可切换查看多CPU模式。
# 键入"z"可切换高亮显示。

# 通过 sar 查看CPU使用率
# 每3秒刷新一次
sar 3
# 每3秒刷新一次，并显示10条
sar 3 10

# 报告有关系统虚拟内存、进程、CPU 活动和 I/O 状态的信息
vmstat 1 10
```





#### 存储

```shell
# 查看当前目录中文件存储占用情况
# [disk usage] 文件或目录
du --block-size=K
# 查看制定文件大小
du --block-size=K beffq Release.md
# 当前
du -sh *

# 磁盘使用情况
# 以“MB”单位查看磁盘占用情况
df --block-size=M
# 前者更简单的方式
df -h
# 查看磁盘剩余空间
df -hl 
```



### 网络

yum 管理包下 netstat 安装的实际包为：`net-tools`

- netstat            监控网络数据包统计数据
- tcpdump        用于网络监控和数据获取的工具



```shell
# 安装网络包
yum install -y net-tools

# 查看linux本地开放端口情况
netstat -luntp
netstat -tulnp

# 前开放的网络连接，检查是否有异常连接
ss -tulwn
```



使用 tcpdump 查看网络请求

```shell
# 查看网卡ens3 (通过 ip addr show) 查看端口 62023 的网络情况
tcpdump -i ens3 port 62023

# 查看任意网卡的 62023 的监听端口
tcpdump -i any port 62023

# 抓包，保存到 http.cap 中，并使用 Wireshark 分析
# -c 20 个包
tcpdump -nni any port 18080 -c 20 -w http.cap
# 持续抓包
tcpdump -nni any port 18080 -w http.cap
```



#### iptables

ip 访问策略

```shell
# 显示当前配置
iptables -L -n

# 显示行号
iptables -nL --line-number
iptables -t nat -nL --line

# 删除第二行
iptables -D INPUT 2

#规则全部清除
iptables -F
```



> 常用策略配置示例

```shell
# 关闭端口[9960]的出入站
iptables -I OUTPUT -ptcp --dport 9960 -j DROP
iptables -I INPUT -ptcp --dport 9960 -j DROP
# 开放80端口外部访问
iptables -I INPUT -p tcp --dport 80 -j ACCEPT

# 允许本地访问
iptables -I INPUT -s 127.0.0.1 -ptcp --dport 9960 -j ACCEPT
iptables -I OUTPUT -s 127.0.0.1 -ptcp --dport 9960 -j ACCEPT
```



#### firewall

防火墙

```shell
# 查看当前活动区域
firewall-cmd --get-active-zones

# 在默认区域放行 TCP 端口 8080
# --permanent 将规则永久保存，即使重启后仍然有效
firewall-cmd --add-port=8080/tcp --permanent
firewall-cmd --add-port=53/udp --permanent
# 移除端口访问策略
firewall-cmd --remove-port=8080/tcp --permanent
# 将源 IP 添加到 internal 区域
firewall-cmd --permanent --zone=internal --add-source=192.168.1.100
# 删除指定的富规则
firewall-cmd --permanent --remove-rich-rule='rule family="ipv4" source address="192.168.1.100" port port="8080" protocol="tcp" accept'


# 重新加载防火墙配置以使更改生效
firewall-cmd --reload

# 验证端口是否已放行
firewall-cmd --list-ports
# 查看特定区域的放行端口
firewall-cmd --zone=public --list-ports
# 查看当前活动区域的所有规则
firewall-cmd --list-all
# 查看富规则
firewall-cmd --list-rich-rules
```



#### ip

> 类似于Windows 的 `ipconfig 命令`

```shell
# 查看当前的网络设备

# 等同于 ipcofnig
ip link

# 等同于 ipconfig
ip addr show

# 显示核心路由表
ip route list
ip route show

# 显示邻居路由表
ip neigh list
ip neigh show
```



#### ssh

ssh 客户端

```shell
# 如使用 ssh 登录 120.63.23.13 服务器
ssh root@120.63.23.13 -p 65222

# 登录成功后查看当前的ip
ifconfig -a
ip addr

# 退出ssh
exit

# scp 文件传输
# 通过scp传输文件
scp /home/sqls/test-database.sql root@10.22.10.231:~

# 下载到本地
scp root@10.22.10.231:~/myfile.txt /home/user/
# 将 /data/bakSql/250108 目录下所有文件下载到当前主机
scp -r root@10.35.2.133:/data/bakSql/250108 ./
# 通配符号下载
scp user@192.168.1.100:/home/user/*.txt /home/localuser/
```



### 其他

#### 日期

主要命令

- date   时间显示
- cal      日历



```shell
# 数据日期
date
# 获取日期并设置格式
date +%Y%m%d-%H%M%S

# 获取秒级时间戳
date +%s

# 获取毫秒级时间戳
date +%s%6N

# 日历显示
cal
```



#### 命令

- type       用于显示命令的类型
- which    显示可执行命令的位置
- whereis   与 which 类似
- man/info  查看命令行的相关帮助信息

```shell
# 显示 cd 命令类型
type cd
# out> type is a shell builtin

type less
# out> less is hashed (/usr/bin/less)

# witch
witch vim
# out> /usr/bin/vim

# 查看命令帮助信息
man vim
vim --help
info vim

# 使用 file 查看文件的类型，看是否为二进制文件
# 如类型："ELF 64-bit LSB executable"或"POSIX shell script"
file $path
```

项目命令的文档信息所在目录：`/usr/share/doc`



##### alias

创建别名(alias)命令，在shell中运行的 alias 仅对当前的session连接有效

```shell
# 创建别名命令
alias a1='cd ~;ll'

# 显示别名命令列表
alias

# 删除命令
unalias a1

# 如创建nginx别名
alias nginx='/usr/local/nginx/sbin/nginx'
whereis nginx
type nginx
```



alias 永久性设置如下：

- 在 `~/.bashrc` 加入 alias命令列表，~表"用户所在目录"
- 运行 `source ~/.bashrc `即可

```shell
# 编辑脚本
vi ~/.bashrc

# 添加 alias 命令列表
source ~/.bashrc
```



**原理**   linux每次启动新的会话，会自动执行~/.bash_profile文件，而这个文件会初始化bashrc。

> source      用于从当前shell会话中的文件读取和执行命令。source命令通常用于保留、更改当前shell中的环境变量。



#### strip 

Removes symbols and sections from files 从文件中删除符号和节。可以建设二进制文件的大小，而不影响其的正常使用。

strip 命令从 XCOFF 对象文件里有选择地除去行号信息、重定位信息、调试段、typchk 段、凝视段、文件头以及全部或部分符号表。 一旦您使用该命令，则非常难调试文件的符号；因此，通常应该仅仅在已经调试和測试过的生成模块上使用 strip 命令。使用 strip 命令降低对象文件所需的存储量开销。

```shell
# 如使用 mingw64 工具删除二进制的全部符号和结
strip -s .\target\release\learning22.exe
```



#### io重定向

将命令行输出到控制台的内容输入指定文件。shell 内部分别将作标准输入、输出和错误称为文件描述符 0、1 和 2。

```shell
# 将 'ls -l ~' 结果输出到文件中
# 使用 '>' 覆盖每次输出的结果到指定的文本中
ls -l ~ > ls-home.txt

# 尾部追加文件内容到文件中
ls -l ~ >> ls-home.txt

# 将错误的结构输出到文件中
ls -1 /dir-not-exist 2> ls-error.txt


# 将标准输出以及错误输出等全部输出到文件中。
# 将标准的文件输出到 ls-output.txt，错误时保持输出一致。
ls -l /bin/usr > ls-output.txt 2>&1
ls -l /bin/usr &> ls-output.txt


# 将错误的输出丢弃掉，'/dev/null' 称 -> 位存储桶
ls -l /bin/usr 2> /dev/null
```



*cat* 命令将多个文件复制到标准输出(stdout)中。

```shell
cat err-*.txt > err-merge.txt

# 改用文件内容输出到控制台。默认为键盘输入
cat < ls-output.txt
```



wc 计子器们，用于统计文件的：行数，字数和字节数。

```shell
# wc 显示：行，字数，字节
ws ls-output.txt

# 统计，bin /usr/bin 中的数量
ls /bin /usr/bin | sort | uniq | wc -l

# 输出含 zip 的列表
ls /bin /usr/bin | sort | uniq | grep zip

# tee 输出 stdin 的数据同时写入到 stdout 和 文件中。
ls /usr/bin | tee ls.txt | grep zip
```





##### 管道线(|)

格式 `commnd1 | command2`  即命令2的输入是命令1的输出。

常用过滤器：sort, uniq

```shell
# 分页显示输出的内容
ll /usr/bin/ | less

# 使用过滤器来先对类别进行排序
ls /usr/bin | sort | less
ls /usr/bin | sort | uniq | less
```



### sudo

> `sudo apt-get install`



> **chmod**  改变目录读写权限 

```shell
# 777  为最高权限
sudo chmod -R 777 <file/direction>

# 755 表示所有者具有读写执行权限 (7)，组成员和其他用户具有读执行权限 (5)。
# 775 所有者具有读写执行权限 (7)，组成员具有读写执行权限 (7)，其他用户具有读执行权限 (5
sudo chmod -R 775 <file/direction>
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

# 查看当前登录的用户
whoami
echo $USER
id
```



### grep 内容搜索



```shell
# 搜索全局文件中的文本
grep -nr "10.35.1.223" *

# 顶级目录下含该文本的所有的
# -n 标识展示行号
grep -nr "10.35.1.223" /*
grep -r "10.35.1.223" /*

# 查找字符串
grep -e str log.txt
# 查找字符串，忽略大小写
grep -i str log.txt
# 查找字符串且显示行
grep -ni str log.txt
```







### 用户与权限

> 用户相关配置

```shell
# 查看系统当前的用户
cat /etc/passwd

# 查看用户的加密形式展示
cat /etc/whadow

# 查看用户分组信息
cat /etc/group

# 查看当前用户
whoami
id
echo $USER

# 密码修改
passwd mysql
# 删除用户密码
passwd -d mysql

# su 切换用户后可在文件中定义环境变量等
vi .bashrc
```



**ACL 权限**

实现多用户对木目录/文件字段的权限管理

```shell
# 查看目录-runtime 的 ALC 权限
getfacl /opt/gyxams/auth/runtime/

# 允许用户-nobody-有读写执行的权限（读/r-read、写/w-write、执行/x-execute权限）
setfacl -R -m u:nobody:rwx /opt/gyxams/auth/runtime/
```





### 硬件

```bash
# 查看当前系统的架构
arch

# 查看系统信息
uname --all

# 获取系统位数
getconf LONG_BIT

# 内存信息
cat /proc/meminfo

# 获取CPU型号信息
cat /proc/cpuinfo
cat /proc/cpuinfo | grep name | cut -f2 -d: |uniq -c

#linux 版本
cat /proc/version

# 查看liunx发新版的信息
cat /etc/*-release
cat /etc/*release*

# 通过hostnamectl查看linux系统信息
hostnamectl

# 查看系统运行时间、负载情况
uptime

# 查看系统加载的内核模块
lsmod

# 查看操作系统版本信息
# 如
cat /etc/system-release
cat /etc/os-release
cat /etc/redhat-release
cat /etc/centos-release
```



### 网络协议

liunx系统支持的协议主要有：telnet, ssh, ftp, sftp, rdp, vnc。


- telnet	    是TCP/IP协议族中的一员，是Internet远程登录服务的标准协议和主要方式	
- ssh			安全外壳协议(secure shell)，最初由 Tatu Ylonen 于 1995 年开发，以取代 Telnet，这是一种允许用户连接到远程计算机的网络协议，通常用于测试连接或远程管理服务器。
- ftp			文件传输协议(file transfer protocol)，用于Internet上的控制文件的双向传输。
- sftp			SSH文件传输协议(SSH File Transfer Protocol/Secret File Transfer Protocol)，安全文件传送协议。
- rdp			远程桌面协议(RDP)是一个多通道(multi-channel)的协议。
- vnc			(Virtual Network Computing)。是基于RFB（Remote Frame Buffer）协议进行通信，是一个基于平台无关的简单显示协议的超级瘦客户系统。



网络测试工具

```shell
# 使用 telnet 测试网络联通通过
telnet 106.32.217.18 3306

# nc/ncat 类似与 tellnet
nc 106.32.217.18 3306
```





### 应用介绍

#### PuTTY

*Windows 下远程连接 Linux 系统，命令行工具。*

*SSH、telnet 客户端，开源项目。*



*开发创始人： Simon Tatham。*



> *终端快捷键（shell/bash等）*

```shell
Ctrl + l(L)   	清屏快捷键，同于 clear 命令
ctrl + w —往回删除一个单词，光标放在最末尾
ctrl + k —往前删除到末尾，光标放在最前面（可以使用ctrl+a）
ctrl + u 删除光标以前的字符
ctrl + k 删除光标以后的字符
ctrl + a 移动光标至的字符头（移动行首）
ctrl + e 移动光标至的字符尾（移动行尾）
ctrl + y 粘贴至光标处
ctrl + r 逆向搜索命令历史（ctrl + g    退出搜索命令）

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



**telnet** 的使用，可用于检查某台服务器的某个端口是否开放或可连通

```shell
# 如连接指定的ip及端口
telent 10.10.16.54 31710
```





#### WinSCP

_*WinSCP*是一个Windows环境下使用SSH的开源图形化SFTP客户端。同时支持SCP协议。它的主要功能就是在本地与远程计算机间安全的复制文件。_

**Linux <=> Windows 之间的文件传输工具(图形化工具)**



软件提供 `PuTTY` 工具集成，通过`ctrl+P`快捷键快速打开；可通过软件保存密码，并且将密码传递给putty，但是软件默认是不开启的。相关配置如下(*版本 5.15.3*)：

- **保存密码** *【新建会话】-> 【保存】中勾选保存密码*
- **传递密码给putty**  *【选项】 -> 【选项】-> 【集成】-> 【应用程序】-> 【记住会话密码并传递给PuTTY(SSH)(R)】*



#### ansiable

一种自动化运维工具,基于paramiko开发的,并且基于模块化工作。由python语言实现。一般用于批量执行任务

```shell
# 批量执行，查看服务器磁盘大小
ansible all -m shell  -a 'fdisk -l| grep "/dev/sd"'
# 查看服务器 cpu 核数
ansible all -m shell  -a 'cat /proc/cpuinfo| grep "processor"| wc -l'
# 查看内存大小
ansible all -a "free -h"

# 查看 host，管理服务器数
cat hosts
```



主要以及常用的配置文件后文件资源

- `/etc/ansible/ansible.cfg`   #配置文件
- `/etc/ansible/hosts`             #主机库（host inventory





## 操作系统

常见的linux发行版操作系统

- Ubuntu
- Fedora
- openSUSE
- Debian
- Red Hat Enterprise Linux (RHEL)
- CentOS            （已停止开发）



国内常见

- openEuler                 兼容 RHEL 包管理
- Anolis OS                  兼容 RHEL 包管理



### Ubuntu

>`dpkg ` 查看安装的软件

```shell
# Ubuntu 分发系统
# 更新系统依赖
sudo apt-get update

# 包搜索
sudo apt search <name>

# 查看系统所有服务列表
service --status-all

# 显示已安装软件列表
dpkg -l
# 添加搜索值
dpkg -l | grep openssl
```





### CentOS

*Community Enterprise Operating System，中文意思是社区企业操作系统*



> 软件 安装/卸载

```shell
# 更新 yum
$ yum update

# yum 软件安装与卸载
$ yum instal <name>
$ yum remove <name>

# 查看 yum 相关的信息
$ yum info <name>

#
# 项目搜索
yum search openldap
# 查看软件包的依赖项
yum deplist openldap

# 查找可安装包的版本号
yum list available openldap.*
yum list available gcc
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

# 启动 php-fpm
/mnt/tool/php/php/sbin/php-fpm
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
cat /etc/passwd | grep -v /sbin/nologin | cut -d : -f 1

# 新增分组 ancient
groupadd ancient

# 添加用户
useradd -g ancient chiyou 		# 将 chiyou 添加到 ancient 分组中
useradd chiyou 

# 查看用户列表
cat /etc/passwd
# 查看用户密码
cat /etc/shadow

# 设置目录数据用户所属
chown -R es:es /home/es/

# 删除用户
userdel -r kingbase

# 查看当前用户分组
group
# 查看 kuafu 用户的所属分组
groups kuafu
# id 用户(查看用户所属组）
id root

# 切换 mysql 到 bash 
sudo -u mysql bash
```



#### rpm

*`Red-Hat Package manager`(same like `RPM Package Manager`)*   RPM软件包管理器，由于其的开放式理念已经使用于许多的Linux发行版。

```shell
# i=> install; v = verbose; h = hash
$ rpm -ivh your-package.rpm
# 更新软件
$ rpm -Uvh soft.version.rpm

# 查看表信息
$ rpm -ql <package>

# 卸载软件
# erase (uninstall) package
$ rpm -e <package>
# --nodeps 前置，无需依赖
$ rpm -e mariadb-libs-5.5.68-1.el7.x86_64 --nodeps

# 更新软件
# upgrade package(s)
$ rpm -U <package>

# 查看应用(libxml2)是否已经安装
$ rpm -qa | grep libxml2
```



- *[rpm 软件搜索及下载地址](http://www.rpmfind.net/linux/RPM/)*
- 地址2 https://pkgs.org/



## 工具包及套件

### 包管理软件

#### dnf

DNF（Dandified Yum）是Fedora Linux操作系统中的一个包管理器，它是基于Yum开发的。它的主要特点是快速、可靠、易用和优秀的用户体验。



### vim

旧版命令 *`vi`*，*`vim`全部兼容前者。`Insert` 或者`i/a`插入文本命令, 推荐使用 `vim`命令。*



> 常用命令

- `:help`			帮助
- `:q`			        退出(quit)
  - `:q!`		取消并退出，不保存
  - `:wq`                退出且保存
  - `ZZ`                  退出且保存
- `:e!`                       丢弃修改并打开原文件



*常用命令*

```shell
# vi/vim 命令
#
#
# 文本查找，向光标之下寻找一个名称为word的字符串
:/search_string
# 向光标之上寻找一个字符串名称为word的字符串
:?search_string
# n 跳的下一个搜索内容（重复前一个搜寻的动作）/next

# windows 与 linux 换行符号不一致的处理方法
:set ff=unix

:显示行号
:set number 
:set nu

# tab 移动
:m,n>  # m到n行右移一个tab
:m,n<  # m到n行左移一个tab
# 多行的话类似，如右移2个tab
:m,n>> 


#
# 跳转至行号
:n
# vi/vim      使用 vi/vim 打开到指定的页面
vi +n filename

# 输出/选择
gg           光标跳转到文件页头   
G            光标跳转到文件页末
ggvG或ggVG   全选高亮
ggyG         全部复制

v 			 进入可视化模式

#选中文本后的执行命令
y			 复制到内容到0寄存器
yy		     复制光标所在的那一行
nyy			 复制光标所在的向下n行
p			 粘贴内容（粘贴至光标后（下））
P			 粘贴至光标前（上）
cc           替换整行，即删除游标所在行并进入插入模式

d			 删除内容
dd			 删除所在行（剪切）
ddp			 交换上下行
x			 上传光标处字符（删除当前字符）
nx			 删除n个连续字符
X			 上传光标处字符（删除前一个字符）
dw/daw       删除单词
d$/D         删除至行尾
d^           删除至行首
dG			 全部删除（删除到文档末尾）
d1G			 删除至文档首部
ndd          删除n行


~			 将光标所在位置字母变成大写或小写
u{n}		 撤销一次或n次操作
U			 撤销当前行的所有修改

>>			 整行将向右缩进
<<			 整行将向左退回

#
# 撤回操作
:u
# 恢复撤销
ctrl+r

# 其他操作
:e			内容重载
:e!			抛弃当前的内容重载

# 文件保存
:w				不退出保存文件。（写入文件）
:w filename 	另保存为其他文件
:w! filename	保存文件并覆盖其他文件

# 设计值行显示
:set nu
# 取消行显示
:set nonu

#对齐操作
:ce			#center/居中
:ri			#right/局右
:le			#left/局左
```



永久显示行数的方法，打开 `vim ~/.vimrc` 输入 `set nu` 保存即可：

```shell
# 打开配置文件
vim ~/.vimrc
# 添加如下内容
set nu
```





> 方位移动

```shell
  k
h   l
  j  
  
  
# 行操作
w,nw	#下一个单词处(word), 也可以实现  nw  多个单词调整
b,nb	#上一个单词处(before word)，支持（nb）
gg      #首行
nG      #跳转值指定的n行
G		#末行
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
- `Ctrl + o`     快速回到上一次光标所在位置
- `ctrl + v`     代码块（文本块）操作
- `shift + i`   进入编辑模式



代码块批操作，如下：*使用 `ctrl + v` 按上下键选择代码块，`shift + i` 进入编辑模式，删除或tab实现批量操作*。

**批量修改**：使用 `ctrl+v` 进入可视模式，在使用 j 等向上下选择，输入 d进行删除；使用 `ctrl+v`进入可视模式，使用 j 等选择块，输入 `#` 等按两次 ESC 即可添加注释。可是模式可使用 ESC 退出



> *搜索与替换*

*在命令模式下，输入 `/` 用于文本查找. 从一般模式进编辑模式，只需按i、I、a、A、o、O、r和R中某个键即可。命令模式：输入：或者/即可进入命令模式。*

*`u`       撤销代码， `Ctrl + R`  恢复撤销*



```shell
# 内容替换
# 进行全局替换
:%s/#/;/g			# 将全部的 "#" 替换为 ";", g 表全部
:1,$s/word1/word2/g或 :%s/word1/word2/g           从第一行到最后一行寻找word1字符串，并将该字符串取代为word2
```







#### 编码

```shell
# vim 字符串乱码的问题

# 查看当前编码格式
echo $LANG

#临时修改编码
LANG="zh_CN.utf8"

# 编码查看
locale -a|grep zh
```



### pacman

可在 [msys2](https://www.msys2.org/) 内使用工具进行包安装，相关命令如下：

```shell
# 对整个系统进行更新
pacman -Syu
# 强制系统更新
pacman -Syy 

# $pkg 包搜索
pacman -Ss <pkg>
# 搜索已安装的包
pacman -Qs <pkg>

# 包安装
pacman -S <pkg>

# gcc 软件包安装
pacman -S --needed base-devel mingw-w64-x86_64-toolchain
```



国内安装包镜像：[清华大学](https://mirror.tuna.tsinghua.edu.cn/help/msys2/)



msys2 供了最新的 GCC，mingw-w64，CPython，CMake，Meson 等原生构建；可以作为 Windows-gun 的环境。



### c library

linux 目前常用的 c library （运行时环境）有这些：

- musl                 linux，轻量级且快速的C库，被设计成简单、一致和正确性优先。
- glibc                 GNU C Library，最广泛使用的标准C库，几乎所有的Linux发行版都默认使用它。
- bionic               android，由Google为Android开发
- uClibc               嵌入式Linux系统设计的C库。它旨在提供一个小型化的C库解决方案，同时保持了足够的功能性和POSIX兼容性，适用于资源受限的环境
- dietlibc             针对嵌入式系统的轻量级C库，目标是创建尽可能小的二进制文件。



查看linux glibc 的版本号

```shell
ldd --version
```



此外其他系统常见的 c library

- msvc                    MSVCRT 是微软Visual C++编译器附带的标准C运行时库。它支持Windows上的应用程序开发，并提供了与标准C和C++兼容的函数实现。
- Newlib                 嵌入式系统、裸机（Bare Metal）、实时操作系统（RTOS）
- Picolibc                嵌入式系统，强调代码尺寸小和性能高
- FreeBSD C Library (libc)     FreeBSD及其他BSD变种



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



*Linux shell 加载脚本如:*

```bash
# shell 脚本加载并执行
source ./real/executable/sellscript.sh
. ./real/executable/sellscript.sh
```



shell脚本开头前缀

```bash
#!/bin/bash
...
...
```



echo 相关语法

```bash
# 输出环境变量
echo $PATH
```





### [bash](http://www.gnu.org/software/bash/manual/)

是 GUN 操作系统的shell 外壳程序，命令行语言解析器。是 *`Bourne-Again SHell`* 的简称，`Bourne shell` 是传统的GUN shell，由 *Stephen Bourne* 编写，前者完全兼容后者。



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



> 流

主要流操作，输出文件 `> file`，输入文件流 `< file`。默认的流是，键盘为输入，显示器(终端)为输出。

```bash
# 输出环境变量到 path.log 文件
echo $PATH > path.log
# 输出到文件，并在文件尾部追加。
echo $PATH >> path.log

# 读取文件
# 用于查看内容较少的纯文本文件。
cat < path.log
# 内容清空
cat /dev/null > test.txt
```





#### 语法

- `#`  为注释符号
- `\`  换行符号
- `''`   单引号，字符串不会执行字符串模板替换，为原始字符串。`'`单引号为原义字符，`"` 转移字符。
- `&&` , `||`  与或运算符号
- `name=[value]` 变量定义，使用`$name` 读取



```bash
# `#` 为注释符号
year=2020

echo '$year is raw string.'
# $year is raw string.
echo "$year is raw string."
# 2020 is raw string.

# 命令
echo "今年是 $(date +%Y) 年，copyright @2018~$(date +%Y)"
```





##### 管道 pipelines

主要使用符号如 `|` 或 `|&`

```shell
# 显示 11 中的第一行
ll | tail -n1
```





##### 语句



> 循环

```bash
#
# 循环
until test-commands;  do consequent-commands; done

while test-commands; do consequent-commands; done

for name [ [in [words ...] ] ; ] do commands; done

for (( expr1 ; expr2 ; expr3 )) ; do commands ; done
# 如
for((i=0;i<100;i++)); do echo $i; done

# break, continue 可提前退出循环
```



> 条件

条件类型或等式比较，https://man7.org/linux/man-pages/man1/test.1.html 。

- string 判别
  - `-n string`          nonzore ，字符串长度非零
  - `-z string`          zore, 字符串长度为零
  - `str = str`          字符串相等
  - `str != str`        字符串不等
- integer 判别
  - `int -eq int`         equal, 整形相等
  - `int -gt int`         greater, 整形大于
  - 其他（水平对比）
    - `-ge`           greater or equal
    - `-le`           less or equal
    - `-lt`           litter (less)
    - `-ne`           not equal
- file 判别
  - 水平对比 `file expor file`
    - `-ef`             equal file，判断两个文件是否相同 是以i节点作为判断的
    - `-nt`             newer than，更新时间新于
    - `-ot`             older than, 更新时间晚于
  - 单判别，`-d file`
    - `-e`               exist，文件存在
    - `-d`               directory，是目录

```shell
# 执行判别并显示结果! 0-成功;1-失败
[ -n '' ] ; echo $?
# 整形比较
[ 1 -eq 1 ] ; echo $?

# 文件比较
[ test.log -ef test.html ]; echo $?
test test.log -ef test.html; echo $?
```





条件语法以及示例如：

```bash
#
# if
if test-commands; then
	consequent-commands;
[elif more-test-commands; then
	more-consequents;]
[else alternate-consequents;]
fi

# 单行 if-elese
if [ 1 -eq 2 ]; then { echo 90;echo true; } else { echo false; } fi


#
# case
case word in
	[ [(] pattern [| pattern]...) command-list ;;]...
esac

#
# 示例；命令直接运行错误，使用脚本则正常
echo -n "Enter the name of an animal: "
# `read` 键盘读取
read ANIMAL
echo -n "The $ANIMAL has "
case $ANIMAL in
    horse | dog | cat) echo -n "four";;
    man | kangaroo ) echo -n "two";;
    *) echo -n "an unknown number of";;
esac
echo " legs."

# select
select name [in words ...]; do commands; done
```



判断应用是否已安装，直接在`shell` 对话框中输出错误。

```shell
#!/bin/bash
echo "请输入软件的名称，如npm："
read name
if ! type ${name} >/dev/null 2>&1;then
        echo "${name} 未安装"
else
        echo "${name} 已安装"
fi
```



#### 快捷键

| 命令   | 描述                     | 其他                                        |
| ------ | ------------------------ | ------------------------------------------- |
| ctrl+A | 把光标移动到命令行开头   |                                             |
| ctrl+E | 把光标移动到命令行结尾。 | 助记：ctrl end                              |
| ctrl+C | 强制终止当前的命令       |                                             |
| ctrl+L | 清屏(clear)              |                                             |
| ctrl+U | 删除或剪切光标之前的命令 |                                             |
| ctrl+K | 删除或剪切光标之后的内容 |                                             |
| ctrl+R | 在历史命令中搜索         | ctrl+research，继续（ctrl+R）转至下一个结果 |
| ctrl+D | 退出当前终端             |                                             |
| ctrl+Z | 暂停，并放入后台         |                                             |
| ctrl+S | 暂停屏幕输出             |                                             |
| ctrl+Q | 恢复屏幕输出             |                                             |



### linux 系统python 版本切换

```shell
# 使用linux系统中版本共享工具（分别定义不同的版本）
sudo update-alternatives --install /usr/bin/python python /usr/bin/python2 100
sudo update-alternatives --install /usr/bin/python python /usr/bin/python3 150

# 进行版本切换
sudo update-alternatives --config python
```



# 工具

## 命令行工具

### curl

```shell
# -v,--verbose   显示详细信息
# curl <url>
# 读取 GitHub 主页信息
curl http://www.httpbin.org/

# 仅显示响应头部信息，可用于body较大时
curl -I http://www.httpbin.org/
# 与 -I 类似
curl --head http://www.httpbin.org/

# 忽略 https 证书错误 --insecure 或 -k
curl -I --insecure 'https://127.0.0.1:7442/'

# 发起 http2，需要 libcurl 支持
curl -I --insecure 'https://127.0.0.1:7442/' --http2

# -X <request-method>
# post 数据信息
curl https://www.httpbin.org/post -X POST -d 'name=Joshua-Conero&age=21' -d 'try=so-many-times'
# 上传数据
curl https://www.httpbin.org/post -F 'imgs=@./qrcode.png' -F 'post_key=12' -F 'k1=s'

# 设置头部
curl -X POST -d 'Username=186******20&Password=L982222' -H 'Content-Type: application/json; charset=utf-8' 'https://www.bignbiot.com/api/account/login'
```



curl 配置 `-d` 默认下会默认添加头部，以及配置请求方式为`POST`。需要json可修改头部 `-H 'Content-Type: application/json'`



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
1. [试试Linux下的ip命令，ifconfig已经过时了 ](https://linux.cn/article-3144-1.html)
1. [Window系统命令行调用控制面板程序](https://www.cnblogs.com/micq/p/15190937.html)
1. [Linux服务器常用巡检命令](https://mp.weixin.qq.com/s/IjHS2l9166mMTt-Vymo3Fg)
1. [msys2 安装 gcc教程](https://blog.csdn.net/ymzhu385/article/details/121449628)



