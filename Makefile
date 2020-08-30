.PHONY: clean
clean:
		mv dbs/test_database/sqls/01_create_table.up.sql dbs/test_database/sqls/up
		mv dbs/test_database/sqls/01_create_table.down.sql dbs/test_database/sqls/down
		rm -f dbs/test_database/sqls/*.sql
		mv dbs/test_database/sqls/up dbs/test_database/sqls/01_create_table.up.sql
		mv dbs/test_database/sqls/down dbs/test_database/sqls/01_create_table.down.sql
		echo "1" > dbs/test_database/version
		rm -f dbs/test_database/schema.sql
		cp schemas/old_schema.sql dbs/test_database/schema.sql
		rm -f db_resolver
		docker-compose down

.PHONY: run
run: main.go cli/*.go cmd/*.go find/*.go util/*.go 
		docker-compose exec db_resolver go build -o db_resolver
		docker-compose exec db_resolver \
		 /go/src/github.com/oshiro3/db_resolver/db_resolver up \
		 test_database \
		 --src gitbucket \
		 --path http://githost:8080/api/v3/repos/root/test/raw/master/schemas/new_schema.sql


.PHONY: up
up:
		docker-compose up -d --build

.PHONY: init
init:
		chmod +x init-mysql.sh
		./init-mysql.sh

.PHONY: test
test:
		
