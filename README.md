# go-web

配置文件

```Ini
[mysql]
host = 10.10.10.10
port = 3306
username = root
password = 123456
database = test
maxIdleConnections = 100
maxOpenConnections = 100
maxConnectionLifeTime = 10
debug = true

```



启动方式
```Bash
# 启动 mysql
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD="123456" -d mysql:8.0.28

# 启动 web服务
$ go run cmd/main.go -c 配置文件的绝对路径

# 命令行请求 或者使用 postman
#{"status":0,"name":"小白","password":"123456","email":"asxxxxxxxxxx@163.com","phone":"xxxxxxxxxxxxx","totalPolicy":0,"hobbySlice":["rap","sing"],"hobby":""}
$ curl -s -XPOST -H'Content-Type: application/json' -d'{"status":0,"name":"小白","password":"123456","email":"asxxxxxxxxxx@163.com","phone":"xxxxxxxxxxxxx","totalPolicy":0,"hobbySlice":["rap","sing"],"hobby":""}' http://127.0.0.1:8080/v1/info/create


```
