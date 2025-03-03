启动
```shell
docker run -d -p 8080:80 \
 --name  monitor-gateway \
 -e ClientID=xxx \
 -e ClientSecret=xxxxx \
 -e Mysql_DbName=dbNmae \
 -e Mysql_Host=127.0.0.1:3306 \
 -e Mysql_IgnoreLog=false \
 -e Mysql_MaxIdleConn=1 \
 -e Mysql_MaxOpenConn=2 \
 -e Mysql_Password=xxx \
 -e Mysql_User=user \
 -e Prom_Addr=http://127.0.0.1:9090 \
 -e Server_HttpPort=80 \
 -e Server_LogLevel="" \
 -e Server_RunMode=debug \
 monitor-gateway:v1.1.0
```