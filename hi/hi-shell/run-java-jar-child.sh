#!/bin/bash

# 查询并杀死旧服务
ps -aux | grep 'hotel-' | grep -v 'grep' | awk '{print $2}' | xargs kill -9

# 遍历当前目录子目录
for i in `ls -d */`; do
    targert_dir=`realpath $i`
    echo "正在执行服务：$i"

    # 目录切换
    cd $i

    # jar 目录
    
    #pwd
    
    jar_name=$(find -maxdepth 1 -name *.jar| head -n 1)
    if [[ -z "$jar_name" ]]; then
        echo "No JAR file found in the directory: $i"
    else
	prj_name=$(basename $(pwd))
        echo "Found JAR file: $jar_name"

        # 使用 nohup 启动 JAR 文件，并将输出重定向到日志文件
	echo "nohup java -Xms128m -Xmx256m -jar $jar_name > $prj_name.log 2>&1 &"
        nohup java -Xms128m -Xmx256m -jar "$jar_name" > $prj_name.log 2>&1 &

        # 输出提示信息
        echo "Started JAR file with nohup. Check myapp.log for details."
    fi


    echo ""
    echo ""
    cd ../
done



