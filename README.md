# databases

Minimal examples of SQLite interaction and MariaDB binlog parsing in Go.

## MariaDB Setup

```
docker pull mariadb:10.5
docker run --name=mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mariadb:10.5 --log-bin=binlog
```
