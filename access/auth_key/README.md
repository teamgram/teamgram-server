# auth_key server
## 为什么要独立出auth_key
> 从安全性考虑，除了auth_key之外，无任何系统可以直接访问auth_key数据库，部署时可以物理隔离，保证数据库的安全

## 功能
- handshake
> 生成auth_key

- query_auth_key
> 提供查询auth_key服务，仅session有权限访问
