package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// "github.com/schemalex/schemalex/diff"
	// 	"io/ioutil"
	"log"
	//"os"
)

const (
	Source   = "file:///go/src/sqls/"
	Database = "mysql://docker:docker@tcp(db:3306)/test_database"
)

func main() {
	sql3, err := ioutil.ReadFile("./schemas/new_schema.sql")
	if err != nil {
	}
	sql4, err2 := ioutil.ReadFile("./schemas/old_schema.sql")
	if err2 != nil {
	}
	f, err3 := os.Create("2_diff_schema.up.sql")
	if err3 != nil {
		log.Printf("read err: %v\n", err3)
	}
	defer f.Close()
	diff.Strings(f, string(sql4), string(sql3), diff.WithTransaction(true))

	f2, err4 := os.Create("2_diff_schema.down.sql")
	if err4 != nil {
		log.Printf("read err: %v\n", err4)
	}
	defer f2.Close()
	diff.Strings(f2, string(sql3), string(sql4), diff.WithTransaction(true))
	// Migrator はコメントや BEGIN COMMIT などの命令を処理できないのでトリミングする必要がある

	m, e := migrate.New("file:///go/src/sqls/", Database)
	if e != nil {
		log.Println("err connection", e)
	}
	log.Printf("%+v\n", m)
	e2 := m.Up()
	// e2 := m.Migrate(2)
	if e2 != nil {
		log.Println("err migrate", e2)
	}
}
