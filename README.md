# databases

I'm not much of a database person, but I felt like playing with SQLite and
MariaDB binlogs.

## MariaDB Setup

```
docker pull mariadb:10.5
docker run --name=mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mariadb:10.5 --log-bin=binlog
```