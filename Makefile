.PHONY: clean
clean:
		rm -f sqls/02_diff_schema*
		rm -f db_resolver
		docker-compose down

.PHONY: run
run: main.go compare.go migrate.go
		docker-compose exec db_resolver go build -o db_resolver
		docker-compose exec db_resolver /go/src/db_resolver

.PHONY: up
up:
		docker-compose up -d --build

.PHONY: init
init:
		chmod +x init-mysql.sh
		./init-mysql.sh
