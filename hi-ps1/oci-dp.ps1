# Oracle 数据导入导出
# 2019年1月9日 星期三


$isExp = Read-Host '数据导出(Y/N)？'
# echo $isExp


# 数据库导出
# expdp system/password@orcl dumpfile=test.dmp schemas=test1
$user = ''
$password = ''
$dbname = ''
$dumpfile = $dbname+'.dmp'
$schemas = $user



# expdp dev/zotion123@mci600a dumpfile=mci600a-dev.dmp schemas=sa
# expdp $user/$password@$dbname dumpfile=$dumpfile schemas=$schemas
echo "expdp $user/$password@$dbname dumpfile=$dumpfile schemas=$schemas"

# 输入任意值结束
# ERROR>> 字符串缺少终止符:
Write-Host 'Press Any Key!' -NoNewline
$null = [Console]::ReadKey('')