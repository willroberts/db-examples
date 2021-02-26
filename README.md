# databases

I'm not much of a database person, but I felt like playing with SQLite, MariaDB
binlogs, and containers.

## MariaDB Setup

```
docker pull mariadb:10.5
docker run --name=mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mariadb:10.5 --log-bin=binlog
```

## MariaDB Binlogs

https://github.com/MariaDB/server/blob/10.6/client/mysqlbinlog.cc
https://github.com/mysql/mysql-server/blob/8.0/client/mysqlbinlog.cc
https://github.com/osheroff/mysql-binlog-connector-java
https://github.com/sysown/proxysql_mysqlbinlog