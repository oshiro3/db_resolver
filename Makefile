.PHONY: clean
clean:
		rm -f sqls/02_diff_schema*
		rm -f differ
		docker-compose down

.PHONY: run
run: main.go compare.go migrate.go
		docker-compose exec diff go build -o differ
		docker-compose exec diff /go/src/differ

.PHONY: up
up:
		docker-compose up -d --build

.PHONY: init
init:
		chmod +x init-mysql.sh
		./init-mysql.sh
