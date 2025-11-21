# 操作系统(OS/Operating System)
>  基本信息
    作者： Joshua Conero
    日期： 2017年6月20日 星期二



## 目录

1. [简介](#menu_overview)
2. [UNIX 1970](#menu_unix)
2. [Windows 1985](./windows.md)
2. [Linux 1991](./linux.md)



# <span id="menu_overview">操作系统通论</span>

## 基本概念

任何计算机系统都包含一个名为操作系统的基本程序集合。该集合内，最重要的程序称为**内核**（kernel）；当操作系统启动时，内核被转入RAM中，内核中包含了系统运行所必不可少的很多核心过程（procedure）。内核也为系统中所有事情提供了主要功能，并决定高层软件的很多特性。因此我们可以将术语“操作系统”等同于“内核”。



操作系统必须完成两个主要目标：

- 与硬件部分交互，为包含在硬件平台上的所有底层可编程部件提供服务
- 未运行在计算机系统上的应用程序（用户程序）提供执行环境



硬件为CPU引入至少两种不同的执行模式：用户程序的非特权模式和内核特权模式，Unix分别将其称为用户态（user mode）和内核态（kernel mode）。



### 进程

所有操作系统都是用一个种基本抽象：进程（process）。一个进程可定义为：“程序执行的一个实例”，或者一个运行程序的“执行上下文”。

操作系统中叫调度程序（scheduler）的部分决定哪个进程能执行。多用户系统中的进程必须是抢占式的（preemptable）；操作系统记录下每个进程占有的CPU时间，并周期性地激活调度程序。



### 内核体系结构

大部分Unix内核都是单块结构：每个内核层都被集成到整个内核程序中，并代表当前进程在内核态下运行。相反，微内核（microkernel）操作系统只需要内核有一个很小的函数集，通常包含几个同步原语、一个简单的调度程序和进程通信机制。



## 启动过程

相关操作文献：**计算机原理**。



启动过程：`Hardware --> BIOS --> BootLoader --> Kernel`。

- （BIOS自检、系统引导）BIOS 检测硬件以及初始化中间表并将BootLoader加载到内存，且控制权移交到BootLoader。
- （启动内核、初始化系统）BootLoader把Kernel写入内存，并执行kernel的 entry point（入口点/初始化程序） 和设置 virtual memory的映射关系。

 



## 类型

当前常见的操作系统有

- Linux
  - Debian
  - Red Hat/Rocky/AlmaLinux
  - SUSE
  - Ubuntu
  - 其他 (generic GLIBC)
- macOS
- Windows
- BSD
  - OpenBSD
  - FreeBSD
  - NetBSD
- Solaris



Linux 是一个从零开始构建的类 Unix 操作系统，而 BSD 则直接继承了 Unix 的血统。





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



