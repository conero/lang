# 版本控制

> Joshua Conero
>
> 2018年8月16日 星期四



*版本控制软件*

# git

> 时间： *2005*

> Staging Area(暂存区)

![](..\image\software\vcs-GitStaging.png)



*分布式版本控制系统（Distributed Version Control System，简称 DVCS） 在这类系统中，像 Git、Mercurial、Bazaar 以及 Darcs 等，客户端并不只提取最新版本的文件快照，而是把代码仓库完整地镜像下来。 这么一来，任何一处协同工作用的服务器发生故障，事后都可以用任何一个镜像出来的本地仓库恢复。 因为每一次的克隆操作，实际上都是一次对代码仓库的完整备份。*



> 部分目标

- 速度
- 简单的设计
- 对非线性开发模式的强力支持（允许成千上万个并行开发的分支）
- 完全分布式
- 有能力高效管理类似 Linux 内核一样的超大规模项目（速度和数据量）



> 基本的 Git 工作流程如下：

1. 在工作目录中修改文件。
2. 暂存文件，将文件的快照放入暂存区域。
3. 提交更新，找到暂存区域的文件，将快照永久性存储到 Git 仓库目录。



## 教程

获取命令行帮助

```console
$ git help config
```

### config

> `git config [options]`  命令行



> 文件位置

1. `/etc/gitconfig` 文件: 包含系统上每一个用户及他们仓库的通用配置。 如果使用带有 `--system` 选项的 `git config` 时，它会从此文件读写配置变量。
2. `~/.gitconfig` 或 `~/.config/git/config` 文件：只针对当前用户。 可以传递 `--global` 选项让 Git 读写此文件。
3. 当前使用仓库的 Git 目录中的 `config` 文件（就是 `.git/config`）：针对该仓库。




> 用户信息

```console
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com
```



**https认证**

> 保存全局 *https* 用户的密码

*git config --global credential.helper store*

> 密码输入错误时，第二次无法输入密码

*git config --system --unset credential.helper*





### Git 基础

> 获取git仓库

- `git init`			初始化git仓库
- `git clone <url>`	从远程仓库中获取git代码仓库



`... @TODO`  [Git-基础-记录每次更新到仓库](https://git-scm.com/book/zh/v2/Git-%E5%9F%BA%E7%A1%80-%E8%AE%B0%E5%BD%95%E6%AF%8F%E6%AC%A1%E6%9B%B4%E6%96%B0%E5%88%B0%E4%BB%93%E5%BA%93)

### 文件管理

#### git 删除未跟踪文件

```ini
#删除 untracked files
git clean -f

# 连 untracked 的目录也一起删掉
git clean -fd

# 连 gitignore 的untrack 文件/目录也一起删掉 （慎用，一般这个是用来删掉编译出来的 .o之类的文件用的）
git clean -xfd

# 在用上述 git clean 前，墙裂建议加上 -n 参数来先看看会删掉哪些文件，防止重要文件被误删
git clean -nxfd
git clean -nf
git clean -nfd

#恢复某个已修改的文件（撤销未提交的修改）
git checkout file-name
git checkout --
git checkout *.php

```

## 主题
### 本地仓库推送多个远程仓库(remote)



1. git remote -v   查看远程仓库
2. git add <name> <url>   添加默认远程仓库，可以新增多个，默认 *origin*
3. git push <name> / <branch> 单个仓库推送
4. git push --all 所有仓库推送数据



# svn

> 非分布式版本控制应用

## 教程

### 分支控制

> 其版本控制，采用目录式的；通常情况下没有默认分支，不指定目录的情况下为全部

- 顶级目录
  - branches		分支
  - tags			标签
  - trunk			主干线



## 参考

### git-scm

> https://git-scm.com/book/zh/v2