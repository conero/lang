# powershell util script
# 2018年11月15日 星期四
# using: $s? or s-help
# powershell >= 6.1


# Szm test equel path
# Using:  s-tEqp -p <path>
function s-tEqp{
	param(
		$p = 'test standard path'
	)
<#
	if([String]::IsNullOrEmpty($p)){
		$p = 'check input is null?'
	}
#>
	# if([String]::IsNullOrEmpty($p)){
		# $p = 'check input is null?'
	# }
	if($p){
		$p = $p.Replace('/', '\')
	}	
	return $p
}

# echo not null value
function s-enn{
	if($args){
		echo "$args"
	}
}

# Szm Add to Path 
# using like: sAddp -p <path> -t <type>
# -p 路径; -t 类型; -n 不执行; -v 检测正确的
function s-Addp{
	param(
		$p,
		$t='User',
		$n=0,
		$v=1
	)
	
	$testMk = 0
	if([String]::IsNullOrEmpty($p)){
		$p = '/any/by/test/C-Joshua'
		$testMk = 1
	}elseif($n){
		$testMk = 1
	}
	
	$op = [System.Environment]::GetEnvironmentVariable("Path", $t)
	$opList = $op -split ";"
	# 数组
	$npList = @()
	$mkCtt = 0
	$p = s-tEqp -p $p
	foreach($o in $opList){
		if(!($p -eq $o)){
			$npList += $o
		}else{
			$mkCtt += 1
		}
	}
	if($mkCtt -eq 0){
		$npList += $p
	}
	$np = $npList -join ";"
	
	if($v){
		$testMk = 1
		if(!(test-path $p)){
			$v = "  ||=> it's a bad value; that can't add to the path."
		}else{
			$v = 0
		}
	}
	# 测试仅仅演示
	if($testMk){
		echo $np
		echo ""
		echo "  ||=> it's just a display..."	
		s-enn $v
	}else{
		[System.Environment]::SetEnvironmentVariable("Path", $np, $t)
		echo '  ||=> 已经设置了环境变量'
	}
}

# 删除错误的路径
# szm del bab path
function s-dbp{
	param(
		$n=0,
		$t='User'
	)
	if($t -eq 'm'){
		$t = 'Machine'
	}
	$op = [System.Environment]::GetEnvironmentVariable('Path', $t)
	$opList = $op -split ';'
	# 数组
	$npList = @()
	$mkCtt = 0
	$p = s-tEqp -p $p
	$i = 0
	$badCtt = 0
	foreach($o in $opList){
		$i += 1
		if(!(test-path $o)){
			# valid in the script file
			echo "${i o}"
			$badCtt += 1
		}else{
			$npList += $o
		}
	}
	
	if($badCtt){
		echo ' `find the bad items: ' + $badCtt
		if(!$n){
			echo ' we`ll update the path'
			$np = $npList -join ';'
			[System.Environment]::SetEnvironmentVariable('Path', $np, $t)
		}
	}else{
		echo ' we scand the $i items, and no bad item is found.'
	}
	
}


# 帮助文档
function s-help{
	echo ' szm - 工具包'
	$list = @(
		's-tEqp', 's-enn', 's-Addp', 's-dbp'
	)
	switch($list){
		's-tEqp'{
			echo ' s-tEqp     szm test equel path'
			echo '	[using] -p <path>'
		}
		's-enn'{
			echo ' s-enn      szm echo not null value'
			echo '	[using] [<any>]'
		}
		's-Addp'{
			echo ' s-Addp     szm add to path'
			echo '	[using] -p <path>'
			echo '	+[using] -t <type>    user(default)/Machine'
			echo '	+[using] -n <no>    	 1/0(default)'
			echo '	+[using] -v <valid>    	 1/0(default);if setted, and -n is setted too'
		}
		's-dbp'{
			echo ' s-dbp      szm del bad path'
			echo '	+[using] -t <type>    user(default)/Machine'
			echo '	+[using] -n <no>    	 1/0(default)'
		}
	}
}

# 别名
$s? = s-help