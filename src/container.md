## Container

> 2024年12月18日
>
> Joshua Conero



容器及编排技术，主要有 docker， kubernetes(k8s)

### Docker

主要核心：容器、镜像、仓库。



工具组成

- Docker Engine            docker 引擎，
  - 对应的github地址：
    - https://github.com/docker-archive/engine    已遗弃
    - https://github.com/moby/moby                      20.10后使用其替代前者
- [docker compose](https://github.com/docker/compose)   用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，您可以使用 YML 文件来配置应用程序需要的所有服务



#### centos 8 安装

> 安装docker

```shell
# 安装docker依赖
yum install -y yum-utils device-mapper-persistent-data lvm2

# 添加阿里云仓库 
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

# 由于 runc 的文件，需要卸载掉 runc 后重新安装 
# 安装 containerd.io
yum install -y http://mirrors.aliyun.com/docker-ce/linux/centos/8.3/x86_64/stable/Packages/containerd.io-1.6.9-3.1.el8.x86_64.rpm

# 安装 docker-ce
yum -y install docker-ce

# 启动服务引擎
systemctl enable --now docker

# 重启docker
#重启daemon进程
systemctl daemon-reload;systemctl restart docker
```



> docker 管理

```shell
# 查看当前的docker信息（配置情况等）
docker info

# 查看docker下服务镜像
docker images

# 服务器重启
docker stop nginx
docker start nginx

docker restart nginx

# 查看服务进程信息
docker ps
# 无截断查看 docker 进程
docker ps --no-trunc
# 查看对应进程使用 compose 执行时的配置文件路径之类的信息，如配置文件地址
docker inspect <container_id_or_name>
```



> docker compose

```shell
# 使用 config 查看配置信息，可用于验证配置文件的正确性
docker-compose config
# 查看服务进程
docker-compose ps

# 启动服务
docker compose up -d
# 强制重启容器
docker compose up -d --force-recreate

# 启动 docker 实例模拟，可用与检查环境等配置信息
docker-compose up --dry-run

# 服务重启
docker-compose restart nginx

# 停止容器，并移除容器
docker compose down
# 停止容器
docker compose stop

# 日志输出跟随（类似 tail -f）
docker composer logs -f
# 指定单个运行的镜像查看日志
docker compose logs digital-smart-gy-conf -f
```



> docker compose 配置文件（docker-compose.yml）

```yaml
services:
  # 列表名称
  web:
    #定义主机名
    container_name: nginx
    image: nginx:latest
    #容器的映射端口
    ports:
      # 将操作系统的8090映射到容器的80端口
      - 8090:80
    #docker 重启后，容器自启动
    restart: always
```



示例如

```yaml
version: "3.8"
services:
  # 系统后端
  imcloud_gms_admin:
    container_name: imcloud_gms_admin
    restart: "always"
    image: openjdk:8-jdk
    hostname: docker_gateway
    network_mode: host
    volumes:
      - ./gms-admin/ruoyi-admin.jar:/root/bin/app.jar
    environment:
      - TZ=Asia/Shanghai
    entrypoint: java -Xms1024m -Xmx1024m  -XX:-UseGCOverheadLimit -Dcsp.sentinel.app.type=1 -jar -Dfile.encoding=utf-8   /root/bin/app.jar #reactor.netty.worker.count=6

  # 系统前端
  imcloud_gms_ui:
    container_name: imcloud_gms_ui
    restart: "always"
    image: nginx
    hostname: docker_gateway
    network_mode: host
    volumes:
      - ./gms-ui/html:/html
      - ./gms-ui/conf/nginx.conf:/etc/nginx/nginx.conf
      - ./gms-ui/logs:/var/log/nginx
      - ./gms-ui/conf.d:/etc/nginx/conf.d
      - ./gms-ui/sslfiles:/var/nginx/sslfiles
    environment:
      - TZ=Asia/Shanghai

```





> docker 容器管理

```shell
# 登陆到容器的bash中
docker exec -it 68b07ea40d4b /bin/bash

# 进行 docker 容器中
docker exec -it ac8d80f6233ec /bin/bash

# 执行，bin/sh
docker exec -it my_container /bin/sh

# 复制文件到容器 
docker cp /opt/app/demo ac8d80f6233ec:/home/demo
# 从容器复制文件
docker cp ac8d80f6233ec:/home/demo /opt/app/demo

# 查看日志 log <container_id&name>
docker logs imcloud_gms_admin
# 以日志追加模式查看，类似于 tail -f
docker logs imcloud_gms_admin -f 
```







### Kubernetes

> @todo

编排技术





### 附录

- docker
  - [docker国内镜像源配置及走代理设置](https://blog.csdn.net/Lichen0196/article/details/137355517)
  - [docker-compose部署nginx，挂载外置配置文件及项目](https://blog.csdn.net/u013652477/article/details/107837931)
  - [Docker-Compose 入门到精通 （图解+秒懂+史上最全）](https://www.cnblogs.com/crazymakercircle/p/15505199.html)