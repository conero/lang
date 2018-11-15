# PowerShell

> 2018年10月19日 星期五

```powershell
# 查看 本地 PS 版本信息
$PSVersionTable
```



## 简介

*Windows PowerShell® 是基于任务的命令行管理程序和脚本语言，专为进行系统管理而设计。*

*UNIX 系统一直有着功能强大的脚本（shell），Windows PowerShell 的诞生就是要提供功能相当于 UNIX 系统 BASH 的命令行外壳程序，同时也内建脚本语言以及辅助脚本程序的工具。*

*PowerShell 使用“谓词 - 名词”命名系统。 每个 cmdlet 名称都由一个标准谓词、连字符和特定名词组成。*



> 设计目标

*PowerShell 旨在消除长期存在的问题和添加新功能，从而改进命令行和脚本环境。*





## 变量

*PowerShell 处理对象。 使用 PowerShell 可以创建称为“变量”的命名对象。 变量名称可以包含下划线字符，也可以是任何字母数字字符。 在 PowerShell 中使用时，始终使用变量名称后跟的 $ 字符指定变量。*



```ini
# 赋值
$jc = 'Joshua Conero is a good guy。'

$jc = $jc + 'We all should be nice to him.'

# Joshua Conero is a good guy。We all should be nice to him.
$jc

# 作为其它变量的别名
$jc = ls
$jc

# 获取环境变量
$jc = Get-ChildItem env:
$jc

# 查看所有变量
Get-Variable
```



> 行号符号

```ini
`			行号符号
```







## 命令参考

```ini
# 获取命令列表
Get-Command

# 查看命令已经分配的别名
Get-Command -CommandType Alias
```



> 关机或重启

```ini
stop-computer
restart-computer
```



### 与 Cmd/ Unix 类似

可在 PowerShell 中使用的常见 cmd.exe 和 UNIX 命令：

|       |         |       |       |
| ----- | ------- | ----- | ----- |
| cat   | dir     | mount | rm    |
| cd    | echo    | move  | rmdir |
| chdir | erase   | popd  | sleep |
| clear | h       | ps    | sort  |
| cls   | history | pushd | tee   |
| copy  | kill    | pwd   | type  |
| del   | lp      | r     | write |
| diff  | ls      | ren   |       |



*PowerShell 别名尝试兼顾清晰度和简洁性。 PowerShell 为常见名词和谓词使用一组标准的别名。*

| 名词或谓词 | 缩写 |
| ---------- | ---- |
| Get        | g    |
| Set        | s    |
| Item       | i    |
| Location   | l    |
| Command    | cm   |
| Alias      | al   |





了解简写名称后，这些别名是可以理解的。

| Cmdlet 名称    | Alias |
| -------------- | ----- |
| `Get-Item`     | gi    |
| `Set-Item`     | si    |
| `Get-Location` | gl    |
| `Set-Location` | sl    |
| `Get-Command`  | gcm   |
| `Get-Alias`    | gal   |





## 参考

- <https://microsoft.com/powershell>



