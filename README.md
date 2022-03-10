# 基于Gin的网关实战

##  一、常用命令

### Golang

```shell
# 初始化项目 go mod
$ go mod init GoHttpServerBestPractice

$ go mod tidy

# golang 打包 【Windows 下编译 Mac 和 Linux 64位可执行程序】
$ SET CGO_ENABLED=0
$ SET GOOS=darwin
$ SET GOARCH=amd64
$ go build -o alias main.go  # go build [-o 输出名] [包名]

$ SET CGO_ENABLED=0
$ SET GOOS=linux
$ SET GOARCH=amd64
$ go build -o alias main.go  # go build [-o 输出名] [包名]
```

### docker

```shell
# docker 设置开机自启动
$ docker update --restart=always [containerId1] [containerId2] [containerId3]

# docke build 
$ docker build -t [tag] .
# 启动
$ docker run -d -v /var/web:/app/static --name=[container-name] -p [8000]:[8080] web

# 删除运行中的容器
$ docker rm -f [containerId]
# 删除停止的容器
$ docker rm [containerId]
# 停止容器
$ docker stop [containerId]
# 删除镜像
$ docker rmi [imageId]
```

### Linux

```shell
# ssh 连接服务器
$ ssh root@192.168.18.100 # 密码 abc123456

# 查看日志文件的最后几行
$ tail -n 5 log.log

# 查看当前所有tcp端口
$ netstat -ntlp 
# $ netstat -ntlp | grep 8080   # 查看所有8080端口使用情况
# $ netstat -ntlp | grep 3306   # 查看所有3306端口使用情况

# 杀掉进程
$ kill 26993 # $ kill PID
# $ kill -9 PID

# 应用在后台执行
$ chmod 777 main
$ nohup ./main > log.log 2>&1.sql &
```

### Git

```shell
# git 取消跟踪文件  https://www.cnblogs.com/zhuchenglin/p/7128383.html
## 对所有文件都取消跟踪
$ git rm -r --cached . 　　// 不删除本地文件
$ git rm -r --f . 　　    // 删除本地文件
## 对某个文件取消跟踪
$ git rm --cached readme1.txt    // 删除readme1.txt的跟踪，并保留在本地
$ git rm --f readme1.txt        // 删除readme1.txt的跟踪，并删除本地文件
## 对某个文件夹取消跟踪
$ git rm --cached log/*    // 删除log文件夹下所有文件的跟踪，并保留文件和文件夹
$ git rm --f log/*         // 删除log文件夹下所有文件的跟踪，并删除文件和文件夹

# 删除远程分支
$ git push origin --delete [branchName]
```

### MySQL

```SQL
# 链接数据库
>>> mysql -h 127.0.0.1.sql -uroot -pabc123456  # mysql -h 127.0.0.1.sql -uroot -pabc123456 -A
# 显示表结构
>>> desc user;
# 显示创建表语句
>>> show create table user;

# mysql导出数据库结构和数据
# 导出整个数据库结构和数据
$ mysqldump -h localhost -P 3306 -uroot -p123456 database > dump.sql
# 导出单个数据表结构和数据
$ mysqldump -h localhost -P 3306 -uroot -p123456  database table > dump.sql
# 导出整个数据库结构（不包含数据）
$ mysqldump -h localhost -P 3306 -uroot -p123456  -d database > dump.sql
# 导出单个数据表结构（不包含数据）
$ mysqldump -h localhost -P 3306 -uroot -p123456  -d database table > dump.sql
# mysqldump 备份导出数据排除某张表
# 就用 --ignore-table=dbname.tablename参数就行了。
$ mysqldump -uusername -ppassword -h192.168.0.1 -P 3306 dbname --ignore-table=dbname.dbtanles > dump.sql
```



## 二、技术选型

| 名称                    | 链接                                               | 安装方式                                      | star  | 说明                    |
| ----------------------- | -------------------------------------------------- | --------------------------------------------- | ----- | ----------------------- |
| viper                   | [链接](https://github.com/spf13/viper)             | go get github.com/spf13/viper                 | 16.1k | golang 配置文件解决方案 |
| go-playground/validator | [链接](https://github.com/go-playground/validator) | go get github.com/go-playground/validator/v10 | 8.9k  | 表单验证                |
| SimonWang00/goeureka    | [链接](https://github.com/SimonWang00/goeureka)    | go get github.com/SimonWang00/goeureka        | 7     | eureka服务注册          |
| gohouse/converter       | [链接](https://github.com/gohouse/converter)       | go get github.com/gohouse/converter           | 219   | 数据库表结构转结构体    |
| gjson                   | [链接](https://github.com/tidwall/gjson)           | go get -u github.com/tidwall/gjson            | 8.8k  | 快速简单的解析json      |
| mapstructure            | [链接](https://github.com/mitchellh/mapstructure)  | go get github.com/mitchellh/mapstructure      | 5k    | map转结构               |
| gin                     | [链接](https://github.com/gin-gonic/gin)           | go get -u github.com/gin-gonic/gin            | 49.4k | web框架                 |
| gorm                    | [链接](https://gorm.io/zh_CN/docs/index.html)      | go get -u gorm.io/gorm                        | 24.4k | orm模型库               |
|                         |                                                    | go get -u gorm.io/driver/mysql                |       | gorm mysql数据库驱动    |
| gotests                 | [链接](https://github.com/cweill/gotests)          | go get -u github.com/cweill/gotests/...       | 3.3k  | 自动生成测试代码        |

## 三、数据库模板

```sql
drop table if exists role;
create table role
(
    id         int unsigned auto_increment primary key,
    created_at int unsigned null comment '创建时间戳',
    updated_at int unsigned null comment '更新时间戳',
    remark     text         null comment '备注',
    deleted_at int unsigned null comment '删除时间戳',
    name       varchar(32)  null comment '角色名',
) ENGINE = InnoDB
  AUTO_INCREMENT = 1.sql
  DEFAULT charset = utf8mb4
  comment ='角色';

drop table if exists user;
create table user
(
    id         int unsigned auto_increment primary key,
    created_at int unsigned null comment '创建时间戳',
    updated_at int unsigned null comment '更新时间戳',
    remark     text         null comment '备注',
    deleted_at int unsigned null comment '删除时间戳',
    name       varchar(64)  null comment '用户名',
    role_id    text         null comment '角色ID: 1.sql,2,3'
) ENGINE = InnoDB
  AUTO_INCREMENT = 1.sql
  DEFAULT charset = utf8mb4
  comment ='用户';
```

