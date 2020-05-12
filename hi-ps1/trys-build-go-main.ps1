#遍历 go 中的main函数项目
#2020年5月12日
#Joshua Conero

$mydir = Read-Host "输入目录，项目循环"
if (-not (Test-Path -Path $mydir)){
    $check = Read-Host "目录无效，是否需要遍历当前目录(y/n)？"
    if ($check -eq "n"){
        exit
    }
    $mydir = "."
}

# 函数定义
function Go-Build-Dirs($dir){
    # 文件遍历
    Get-ChildItem $dir | foreach{
        $vpath = "$dir/" + $_.name
        #echo $vpath, (Test-Path $vpath -PathType Container)

        # 目录
        if (Test-Path $vpath -PathType Container){
            # 文件查找
            $mainFile2 = "$vpath/main.go"
            if (Test-Path $mainFile2 -PathType Leaf){
                echo "$mainFile2 正在编译..."
                go build $mainFile2
                continue
            }

            echo $vpath
            Go-Build-Dirs $vpath
        }



        # 文件查找
        $mainFile = "$dir/main.go"
        echo $mainFile
        if (Test-Path $mainFile -PathType Leaf){
            echo "$mainFile 正在编译..."
            go build $mainFile
            break
        }
    }
}

Go-Build-Dirs $mydir

Read-Host "enter any keys"

