#更新 svn 且合并到 git 日志中

#获取当前 svn 的版本号
$rversion = svn info --show-item revision;

#svn 更新
svn update

# 查看状态
git status
$need = Read-Host "[r$rversion] 您需要合并操作吗（y/n）"
if ('n' -eq $need){
    Read-Host "输入任何字符退出"
    exit
}

# git 提交
git add .
git commit -m "svn: svn-merge. ps1-auto(r$rversion)"
#>

Read-Host "输入任何字符退出"