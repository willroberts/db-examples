.PHONY: mariadb
mariadb:
	docker pull mariadb:10.5
	docker run --name=mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mariadb:10.5 --log-bin=binlog
.PHONY: mysql_bootstrap
mysql_bootstrap:
	@ $(eval ROOT_PASSWORD = $(shell cat .root-password))
	@ echo 'Initializing MySQL container'
	@ docker run --name=mysql -p 3306:3306 -d \
		-e MYSQL_ROOT_PASSWORD=${ROOT_PASSWORD} \
		mysql:8 && echo 'MySQL container started'

.PHONY: mysql_start
mysql_start:
	@ docker start mysql && echo 'MySQL container started'

.PHONY: mysql_stop
mysql_stop:
	@ docker stop mysql && echo 'MySQL container stopped'

.PHONY: schema
schema:
	@ $(eval ROOT_PASSWORD = $(shell cat .root-password))
	@ echo 'Creating databases'
	@ mysql -uroot -p${ROOT_PASSWORD} -h127.0.0.1 < schema/databases.sql
	@ echo 'Creating tables'
	@ mysql -uroot -p${ROOT_PASSWORD} -h127.0.0.1 test < schema/tables.sql
	@ echo 'Creating indexes'
	@ mysql -uroot -p${ROOT_PASSWORD} -h127.0.0.1 test < schema/indexes.sql
	@ echo 'Creating test user'
	@ go run cmd/create_user/main.go
	@ echo 'Populating test data'
	@ mysql -uroot -p${ROOT_PASSWORD} -h127.0.0.1 test < schema/testdata.sql
