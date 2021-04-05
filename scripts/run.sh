#!/usr/bin/env bash
#Remove all container
OUTPUT=$(sudo docker container ls -aq )
echo "$OUTPUT" | while read -r a; do docker container rm $a; done

#Run etcd
docker run --name etcd-docker -d -p 2379:2379 -p 2380:2380 appcelerator/etcd
#Run mysql
docker run --name mysql-docker -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql:5.7
#Run redis
docker run --name redis-docker -p 6379:6379 -d redis 
docker start redis-docker mysql-docker etcd-docker

sleep 10
docker exec -it mysql-docker sh -c 'exec mysql -u root -p -e"CREATE DATABASE chatengine;"' 
docker exec -i mysql-docker mysql --user=root chatengine < /home/anhttn/go/src/github.com/nebula-chat/chatengine/scripts/chatengine.sql
docker exec -i mysql-docker mysql --user=root chatengine < /home/anhttn/go/src/github.com/nebula-chat/chatengine/scripts/chatengine-fix.sql
docker exec -i mysql-docker mysql --user=root chatengine < /home/anhttn/go/src/github.com/nebula-chat/chatengine/scripts/merge-20181129.sql
docker exec -i mysql-docker mysql --user=root chatengine < /home/anhttn/go/src/github.com/nebula-chat/chatengine/scripts/merge-20181214.sql
docker exec -i mysql-docker mysql --user=root chatengine < /home/anhttn/go/src/github.com/nebula-chat/chatengine/scripts/merge-20181220.sql
docker exec -i mysql-docker mysql --user=root chatengine < /home/anhttn/go/src/github.com/nebula-chat/chatengine/scripts/merge-20190619.sql
docker exec -i mysql-docker mysql --user=root chatengine < /home/anhttn/go/src/github.com/nebula-chat/chatengine/scripts/merge-react.sql
