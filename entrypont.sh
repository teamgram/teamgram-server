#!/bin/bash

if [ ! -n "$CHATENGINE_HOST" ]; then
  echo ">>> Plase set environment variable CHATENGINE_HOST to your own server IP. <<<"
  exit 1
fi

export CHATENGINE_HOST=${CHATENGINE_HOST}
export ETCD_URL=${ETCD_URL:-"http://etcd:2379"}
export REDIS_HOST=${REDIS_HOST:-"redis:6379"}
export MYSQL_URI=${MYSQL_URI:-"root:chatengine@tcp(mysql:3306)/chatengine?charset=utf8mb4"}

# create configs from config templates.
createConfigs() {
  CONFIG_TARGET_DIR=/app
  CONFIG_TEMPLATES_DIR=/app/config-templates
  for file in `ls $CONFIG_TEMPLATES_DIR`; do
    cat $CONFIG_TEMPLATES_DIR/$file \
      | sed 's#"ip_address": "192.168.1.150"#"ip_address": "'$CHATENGINE_HOST'"#g' \
      | sed "s#http://127.0.0.1:2379#$ETCD_URL#g" \
      | sed "s#127.0.0.1:6379#$REDIS_HOST#g" \
      | sed "s#root:@tcp(127.0.0.1:3306)/chatengine?charset=utf8mb4#$MYSQL_URI#g" \
      | cat > $CONFIG_TARGET_DIR/$file
  done
}

starService() {
    echo "starting $1/$2..."
    ./$2 >> /tmp/$1_$2.log &
    echo "running $1/$2!"
}

createConfigs

starService service document
starService service auth_session
starService messenger sync
starService messenger upload
starService messenger biz_server
starService access auth_key
starService access session
starService access frontend

tail -f /dev/null