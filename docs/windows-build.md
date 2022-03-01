# 编译运行nebula的开源服务端chatengine笔记

> 经作者[@robinfoxnan](https://me.csdn.net/robinfoxnan) 同意转载
> 

官方说在LINUX下运行，但是看了一下，基本用纯GO来编译，应该可以在WINDOWS下运行。

我使用的Windows10家庭版本，无法使用docker，所以决定直接安装服务。

对于telegramd，按照作者的说法是不再支持，所以建议大家使用chatengine。

## 1. 安装mysql5.7.19

　[下载地址：](https://dev.mysql.com/downloads/mysql/)

用自带的执行脚本。\chatengine\docker\mysql\init-sql\01-chatengine.sql

需要注意：我是使用navicat进行建库和导入表结构，低版本不支持数据类型，而5.7安装比较麻烦，可以参考我另一片文章。

其中5.7不支持脚本中的时间戳默认的值，需要在会话中加一个选项：

```
set session sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
```

否则会报告很多错误，造成表丢失，后期运行时候爆错。

## 2. 安装redis

[帮助：](https://www.cnblogs.com/jylee/p/9844965.html)

[下载地址：](https://github.com/MicrosoftArchive/redis/releases)

这个没有太多说的，比较简单。

## 3. 安装etcd

[介绍：](https://blog.csdn.net/skh2015java/article/details/80712214)

[下载地址：](https://github.com/etcd-io/etcd/releases)

需要添加一个配置文件：命名为：etcd.conf.yml

```
name: etcd
listen-client-urls: http://0.0.0.0:2379
advertise-client-urls: http://0.0.0.0:2379
```

启动时候执行命令：

```
etcd.exe  –config-file etcd.conf.yml > log.txt
```
 
我使用golang1.13.8

在c:\godir   设置为GOPATH

注意：目录一定要按照约定方式存放，比如

C:\godir\src\github.com\nebula-chat\chatengine

否则编译时候会找不到库，

按照手册逐个的go get / go build

然后将各个配置文件的地址更改一下。

 

启动各个文件比较繁琐，所以做一个批处理文件用来启动各个EXE。

 
## 注意
注意：目前开源的版本支持1.0协议，所以编译客户端时候需要按nubela说明，checkout到指定版本打补丁编译。

 支持group，但是不支持channel和secret聊天。

企业版（收费版）支持目前的新的功能，有需要的自己telegram联系作者。


具体源码的分析，见我其他的文章。
