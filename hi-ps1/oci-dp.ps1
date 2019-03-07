# Oracle 数据导入导出
# 2019年1月9日 星期三
# 参考 link+ https://blog.csdn.net/jiandonghan/article/details/81259093


$isExp = Read-Host '数据导出(Y/N)？'
# echo $isExp


# 数据库导出
# expdp system/password@orcl dumpfile=test.dmp schemas=test1
$user = ''
$password = ''
$dbname = ''
$dumpfile = $dbname+'.dmp'
$schemas = $user


# 数据导入
# test1 => test2; 前面用户导入到后面的
# impdp system/password@orcl dumpfile=test.dmp remap_schema=test1:test2
# dmp 文件: ~/admin/dbname/dpdump

# expdp dev/zotion123@mci600a dumpfile=mci600a-dev.dmp schemas=sa
# expdp $user/$password@$dbname dumpfile=$dumpfile schemas=$schemas
echo "expdp $user/$password@$dbname dumpfile=$dumpfile schemas=$schemas"

# 输入任意值结束
# ERROR>> 字符串缺少终止符:
Write-Host 'Press Any Key!' -NoNewline
$null = [Console]::ReadKey('')