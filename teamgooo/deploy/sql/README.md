# init database

## **推荐复位方式**
先停服务，再同时复位 MySQL 和 Redis，最后重启并让客户端重新登录。

```bash
cd /opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/teamgooo/bin
./killall.sh
```

先 dry-run 看会清哪些表：

```bash
cd /opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2

TEAMGOOO_RESET_ENV=local \
TEAMGOOO_RESET_DSN='root:@tcp(127.0.0.1:3306)/teamgooo' \
bash teamgooo/deploy/sql/reset_non_users_for_int64_time.sh --dry-run
```

确认后执行 DB 复位：

```bash
cd /opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2

TEAMGOOO_RESET_ENV=local \
TEAMGOOO_RESET_DSN='root:@tcp(127.0.0.1:3306)/teamgooo' \
TEAMGOOO_RESET_CONFIRM_DB=teamgooo \
bash teamgooo/deploy/sql/reset_non_users_for_int64_time.sh
```

这个脚本会保留 `users` 表，清掉其他表，包括 `auth_keys` / `auth_users` / 消息链路状态表，并重新初始化 `userupdates_partition_fences`。

**Redis 必须同步清**
如果本机 Redis 只给这个 Teamgooo 环境用，最干净：

```bash
redis-cli -h 127.0.0.1 -p 6379 FLUSHDB
```

如果 Redis 不是独占的，就至少清 Teamgooo 相关 key：

```bash
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'authsession:auth_data:v1:*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'auth_key_ttl#*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'salts#*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'idgen:malloc_seq:*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'user_data.v3#*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'user:facts:v1:*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'user:privacy:v1:*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'user:contact-map:v1:*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
redis-cli -h 127.0.0.1 -p 6379 --scan --pattern 'user:presence:v1:*' | xargs -r redis-cli -h 127.0.0.1 -p 6379 DEL
```

然后重启服务：

```bash
cd /opt/data/teamgram2/src/github.com/teamgram/teamgram-server-v2/teamgooo/bin
./r.sh
```

复位后客户端必须重新登录。否则旧连接/旧 auth cache 和新 DB 状态不一致，还是会出现 push 找不到 auth key、消息显示异常的问题。