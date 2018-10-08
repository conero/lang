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



*Git 保存的不是文件的变化或者差异，而是一系列不同时刻的文件快照。*

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



> *保存用户密码*

*credential cache* 

`git config --global credential.helper cache`





**https认证**

> 保存全局 *https* 用户的密码

*git config --global credential.helper store*

> 密码输入错误时，第二次无法输入密码

*git config --system --unset credential.helper*





### Git 基础

#### 获取git仓库

- `git init`			初始化git仓库
- `git clone <url>`	从远程仓库中获取git代码仓库



#### 记录每次更新到仓库

> **已跟踪或未跟踪**



*你工作目录下的每一个文件都不外乎这两种状态：已跟踪或未跟踪。 已跟踪的文件是指那些被纳入了版本控制的文件，在上一次快照中有它们的记录，在工作一段时间后，它们的状态可能处于未修改，已修改或已放入暂存区。 工作目录中除已跟踪文件以外的所有其它文件都属于未跟踪文件，它们既不存在于上次快照的记录中，也没有放入暂存区。 初次克隆某个仓库的时候，工作目录中的所有文件都属于已跟踪文件，并处于未修改状态。*



> **文件的状态变化周期**

![文件的状态变化周期](../image/software/vcs/lifecycle.png)



> **`$ git status` 查看本次仓库的状态**

*状态简览* `# git status -s|--short`





> **`$ git add <file>` 跟踪文件，即添加文件到暂存区**



>  **忽略文件 `.gitignore` 可共享策略**

文件 `.gitignore` 的格式规范如下：

- 所有空行或者以 `＃` 开头的行都会被 Git 忽略。
- 可以使用标准的 glob 模式匹配。
- 匹配模式可以以（`/`）开头防止递归。
- 匹配模式可以以（`/`）结尾指定目录。
- 要忽略指定模式以外的文件或目录，可以在模式前加上惊叹号（`!`）取反。



> **`$ git diff`   此命令比较的是工作目录中当前文件和暂存区域快照之间的差异， 也就是修改之后还没有暂存起来的变化内容。**



*`$ git diff --check`  找到可能的空白错误并将它们为你列出来*



`$ git diff --cached`   或 `$ git diff --staged`	*查看已暂存的将要添加到下次提交里的内容 (--staged 和 --cached 是同义词）*



> **`$ git commit` 提交更新**

提交之前可以`git status` 查看状态

`$ git commit -m "massage"`   *将提交信息与命令放在同一行。 原则:  一般情况下，信息应当以少于 50 个字符（25个汉字）的单行开始且简要地描述变更，接着是一个空白行，再接着是一个更详细的解释. Git 项目要求一个更详细的解释，包括做改动的动机和它的实现与之前行为的对比 - 这是一个值得遵循的好规则。*





**跳过使用暂存区域**

`$ git commit -a`   *Git 就会自动把所有已经跟踪过的文件暂存起来一并提交，从而跳过 `git add` 步骤*



> **`$ git rm <file>`  移除文件, <file>  可以为`glob` 模式**

*从暂存区移除文件*



- `git rm -f`   *使用强制(强制删除)删除*
- `git rm --cached`  *从暂存区删除文件，但是保持在工作目录中*



> **`$ git mv <source-file> <target-file>` *移动文件***



```console
$ git mv README.md README
$ git status
On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

    renamed:    README.md -> README
```

运行 `git mv` 就相当于运行了下面三条命令：

```console
$ mv README.md README
$ git rm README.md
$ git add README
```



#### 查看提交历史

> `$ git log`

- `git log -p -2` *一个常用的选项是 `-p`，用来显示每次提交的内容差异。 你也可以加上 `-2` 来仅显示最近两次提交*

- `git log --stat`  *如果你想看到每次提交的简略的统计信息，你可以使用 `--stat` 选项*



*另外一个常用的选项是 `--pretty`。 这个选项可以指定使用不同于默认格式的方式展示提交历史。 这个选项有一些内建的子选项供你使用。 比如用 `oneline` 将每个提交放在一行显示，查看的提交数很大时非常有用。 另外还有 `short`，`full` 和 `fuller` 可以用，展示的信息或多或少有些不同，请自己动手实践一下看看效果如何。*



#### 撤消操作

- `$ git commit --amend`    *提交有问题是时，重新提交*



> **`$ git reset HEAD <file>` *取消暂存的文件***

*在调用时加上 `--hard` 选项**可以**令 `git reset` 成为一个危险的命令*



> **`$ git checkout -- <file>...` *撤消对文件的修改***

*你需要知道 `git checkout -- [file]` 是一个危险的命令，这很重要。*



#### 远程仓库

> **`$ git remote -v` *查看远程仓库*** 



> **`$ git remote add <shortname> <url>` *添加远程仓库*** 



> **`$ git fetch [remote-name]` *从远程仓库中抓取与拉取***



> **`$ git remote show origin` *查看远程仓库***



> **`$ git remote rename <old-name> <new-name>` *远程仓库的移除与重命名***



#### 打标签

> `$ git tag `



> **`$ git tag` *列出标签时***

`$ git tag -l v1.8.5*` 使用 *glob* 模式筛选



> **创建标签**

*Git 使用两种主要类型的标签：轻量标签（lightweight）与附注标签（annotated）。*

`$ git tag <tag-name> `

`$ git tag -a v1.4 -m 'my version 1.4'`   *附注标签, 创建标签时添加标签*



> **共享标签**

`$ git push origin <tag-name>` *推送标签到远程仓库*

`$ git push origin --tags`  *一次性推送很多标签*







### 分支

> Git 鼓励在工作流程中频繁地使用分支与合并



`提交对象(commit object)` *包含信息 :*

- 指向暂存内容快照的指针
- 作者的姓名和邮箱
- 提交时输入的信息
- 指向它的父对象的指针

*首次提交产生的提交对象没有父对象，普通提交操作产生的提交对象有一个父对象，而由多个分支合并产生的提交对象有多个父对象*



*暂存操作会为每一个文件计算校验和（使用 SHA-1 哈希算法），然后会把当前版本的文件快照保存到 Git 仓库中（Git 使用 blob 对象来保存它们），最终将校验和加入到暂存区域等待提交*



*Git 的分支，其实本质上仅仅是指向提交对象的可变指针。 Git 的默认分支名字是 `master`。 在多次提交操作之后，你其实已经有一个指向最后那个提交对象的 `master` 分支。 它会在每次的提交操作中自动向前移动。*



*Git 的 “master” 分支并不是一个特殊分支。 它就跟其它分支完全没有区别。 之所以几乎每一个仓库都有 master 分支，是因为 `git init` 命令默认创建它，并且大多数人都懒得去改动它。*



`分支创建`    *创建了一个可以移动的新的指针*

`HEAD` 	    *是一个指针，指向当前所在的本地分支（译注：将 `HEAD` 想象为当前分支的别名）。 HEAD 分支随着提交操作自动向前移动*



#### 新增分支

```ini
# 拉取分支，若不存在就新建分支
# Method 1
$ git checkout -b <branchName>

# 分支简写
# Method 2
$ git branch <branchName>
$ git checkout <branchName> 
```



#### 切换分支

```ini
# 切换分支，回到主线分支
git checkout <branchName>
```



#### 分支合并

```ini
# 转到需要合并的分支(目标分支)
$ git checkout <targetBN>

# 与源分支合并
$ git merge <srcBN>
```



*合并提交，它的特别之处在于他有不止一个父提交。*



#### 删除分支

```ini
$ git branch -d <branchName>
```



> *长期分支*, *平行分支*， *特性分支*

*平行分支*     `master <->  develop/next`



#### 推送远程

`$ git push origin <remoteBranchName>`  推送到远程仓库

`$ git push origin local:remote`	推送:   本地 -> 远程



#### 跟踪分支

*从一个远程跟踪分支检出一个本地分支会自动创建一个叫做 “跟踪分支”（有时候也叫做 “上游分支”）。跟踪分支是与远程分支有直接关系的本地分支。 如果在一个跟踪分支上输入 `git pull`，Git 能自动地识别去哪个服务器上抓取、合并到哪个分支。*



`$ git checkout -b [branch] [remotename]/[branch]`

`$ git checkout --track origin/serverfix`



```ini
# 查看所有跟踪的远程分支
$ git branch -vv

# 查看所有分支
$ git fetch --all; git branch -vv
```



#### 变基

*在 Git 中整合来自不同分支的修改主要有两种方法：`merge` 以及 `rebase`。*



###  Git 服务器

**裸仓库**  *一个远程仓库通常只是一个裸仓库（*bare repository*）— 即一个没有当前工作目录的仓库。 因为该仓库仅仅作为合作媒介，不需要从磁碟检查快照；存放的只有 Git 的资料。 简单的说，裸仓库就是你工程目录内的 `.git` 子目录内容，不包含其他资料。*

`$ git clone --bare <urlName>`     获取裸仓库(以裸仓库克隆服务器)



> 协议

*本地协议（Local），HTTP 协议，SSH（Secure Shell）协议及 Git 协议*



- `ssh`   *授权访问*

- `git`    *无授权访问*

- `Smart HTTP`   访问上同时支持 *授权/无授权* 协议

​	*主要原理是使用一个 Git 附带的，名为 `git-http-backend` 的 CGI。它被引用来处理协商通过 HTTP 发送和接收的数据。 它本身并不包含任何授权功能，但是授权功能可以在 Web 服务器层引用它时被轻松实现。*



*Git 提供了一个叫做 GitWeb 的 CGI 脚本来实现基于网页的简易查看器*



### 分布式 Git 

> 分布式工作流

*集中式工作流、集成管理者工作流、司令官与副官工作流*



### github

> fork

*派生（Fork）项目:  GitHub 将在你的空间中创建一个完全属于你的项目副本，且你对其具有推送权限*



> 合并请求（Pull Request）





### Git 工具







`... @TODO`  [7.1 Git 工具 - 选择修订版本](https://git-scm.com/book/zh/v2/Git-%E5%B7%A5%E5%85%B7-%E9%80%89%E6%8B%A9%E4%BF%AE%E8%AE%A2%E7%89%88%E6%9C%AC)





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



## 附录

### 中英文对照

| 英文          | 中文                      | 说明 |
| :------------ | :------------------------ | :--- |
| track/untrack | 跟踪(文件)/未跟踪的(文件) |      |
| stage         | 暂存                      |      |
| commit        | 提交                      |      |
| repository    | 仓库                      |      |
|               |                           |      |
|               |                           |      |
|               |                           |      |
|               |                           |      |
|               |                           |      |



> 其他词汇

- `commit object`  提交对象
- `decorate`      

### 参考

#### git-scm

> https://git-scm.com/book/zh/v2



# svn

> 非分布式版本控制应用

## 教程

### 分支控制

> 其版本控制，采用目录式的；通常情况下没有默认分支，不指定目录的情况下为全部

- 顶级目录
  - branches		分支
  - tags			标签
  - trunk			主干线
