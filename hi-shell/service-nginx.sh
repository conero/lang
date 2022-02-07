#!/bin/sh
# --> 文件命名: /etc/init.d/nginx
# linux nginx 服务脚本: 2021年11月18日 星期四
# 语法: service [serviceName] [signal]
# service nginx start


# 服务名称
SERVERNAME="nginx"
DAEMON=/usr/local/nginx/sbin/nginx
# UUIDGEN=/usr/bin/dbus-uuidgen
# UUIDGEN_OPTS=--ensure
NAME=nginx
# DAEMONUSER=messagebus
# PIDDIR=/var/run/dbus
# PIDFILE=$PIDDIR/pid
DESC="nginx 服务器"

# 服务启动
start()
{
    echo "服务启动中……"
    /usr/local/nginx/sbin/nginx
    exit 0;
}

# 服务停止
stop()
{
    echo "服务停止中……"
    /usr/local/nginx/sbin/nginx -s stop
    exit 0;
}

reload()
{
    echo "服务重载中……"
    /usr/local/nginx/sbin/nginx -s reload
    exit 0;
}

case "$1" in
start)
    start
    ;;
stop)
    stop
    ;;
reload)
    reload
    ;;    
restart)
    stop
    start
    ;;
*)
    echo "usage: $0 start|stop|restart|reload"
    exit 0;
esac
exit