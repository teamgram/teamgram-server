# upload服务
> telegram客户端实现里，upload和download是两个独立的连接，相应的会有upload和download独立的服务。
>
> 当前先简单实现upload和download逻辑
>
> 后续也将upload和download服务独立出来，并引入独立的图片存储服务和文件存储系统

## TODO
未使用分布式文件系统

## 多机器部署注意事项
upload多机部署时，需要使用nfs，将/opt/nbfs目录mount到本机
