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
    echo -e "\033[35m服务启动中…… \033[0m"
    /usr/local/nginx/sbin/nginx
    echo -e "\033[32m [SUCCESS] \033[0m 服务已启动！"
}

# 服务停止
stop()
{
    echo "\033[35m服务停止中…… \033[0m"
    /usr/local/nginx/sbin/nginx -s stop
    echo -e "\033[32m [SUCCESS] \033[0m 服务已停止！"
}

reload()
{
    echo "\033[35m服务重载中…… \033[0m"
    /usr/local/nginx/sbin/nginx -s reload
    echo -e "\033[32m [SUCCESS] \033[0m 服务已重载配置文件！"
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
    echo -e "\033[36m usage: $0 start|stop|restart|reload \033[0m"
esac
exit 0;