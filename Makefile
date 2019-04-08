.PHONY: setup init history status up e2e mysql.start mysql.stop

setup:
	@if [ -z `which sql-migrate 2> /dev/null` ]; then \
		go get github.com/rubenv/sql-migrate/...; \
	fi

init:
	mysql -h 127.0.0.1 --port $(MYSQL_PORT) -u$(MYSQL_USER) < db/database.sql

status: setup
	sql-migrate status

up: setup
	sql-migrate up

hisotry:
	mysql -h 127.0.0.1 --port $(MYSQL_PORT) -u$(MYSQL_USER) -D caww -e "SELECT * FROM gorp_migrations;"

e2e:
	go test -tags e2e ./...

mysql.start:
	docker run --rm -d -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p $(MYSQL_PORT):3306 --name mysql_caww mysql:5.7

mysql.stop:
	docker stop mysql_caww
