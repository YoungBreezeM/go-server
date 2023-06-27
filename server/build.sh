#! /bin/bash
basepath=$(pwd)
mkdir -p build && cd build
touch config.yaml

cat > config.yaml <<EOF
# server
host: 127.0.0.1
port: 8080

#level
log:
  level: 5

#mysql
mysql: 
  host: localhost
  username: root
  port: 3306
  password: root
  level: 4

# wechat
wechat:
  appId: xxxxxxxxxxxxxxx
  appSecert: xxxxxxxxxxxxxxxxxxxxxx
  token: xxxxxxxxxxxxxxxxxxxxxxxxxx
EOF
go build -o server -v -ldflags="-s -w" $basepath/main.go