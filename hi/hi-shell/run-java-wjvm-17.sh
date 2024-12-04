#!/bin/bash

# 服务器运行
java17=/usr/local/java17/bin/java

# 设置环境变量
export JAVA_HOME=/usr/local/java17
# 查看 java 版本好用于验证
$java17 -version

# 声明函数
# function start() { # function 不兼容部分 linux 系统
start() {
    # 启动服务
    # nohup $java17 -Xms256m -Xmx518m -jar /home/blockmg2412/bmg-server/block-admin.jar --spring.config.location=/home/blockmg2412/bmg-server/application.yml &
    nohup $java17 -Xms256m -Xmx518m -jar /home/blockmg2412/bmg-server/block-admin.jar &
}

stop() {
    # 停止服务
    kill -9 `ps -ef|grep java|grep -v grep|awk '{print $2}'`
}

# 判断服务是否在运
status() {
    # 判断服务是否在运
    pid=`ps -ef|grep java|grep -v grep|awk '{print $2}'`
    if [ -n "$pid" ]; then
        echo "服务已启动，pid: $pid"
    else
        echo "服务未启动"
    fi
}

# 启动服务
case $1 in
    start)
        start
        ;;
    stop)
        stop
        ;;
    status)
        status
        ;;
    *)
        echo "Usage: $1 {start|stop|status}"
        exit 1
        ;;
esac