# mysql_condition

## Usage
1回目は失敗するが原因は SQL shema に余分な内容が入ってくるから。

それをトリムして diff 作成部分をコメントアウトして再度実行すると上手くいく。

```
$ sudo docker-compose up -d
$ sudo docker-compose exec diff bash
# cd /go/src
# go build
# go run main.go
```
