# 版本控制

> Joshua Conero
>
> 2018年8月16日 星期四



*版本控制软件*

# git
## 教程
### config

> 保存全局 *https* 用户的密码

*git config --global credential.helper store*

> 密码输入错误时，第二次无法输入密码

*git config --system --unset credential.helper*





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

