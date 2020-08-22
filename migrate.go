package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

const (
	Source   = "file:///go/src/sqls/"
	Database = "mysql://docker:docker@tcp(db:3306)/test_database"
)

func schemaMigrate() {

	m, e := migrate.New(Source, Database)
	if e != nil {
		log.Println("err connection", e)
	}
	// e2 := m.Up()
	e2 := m.Migrate(2)
	if e2 != nil {
		log.Println("err migrate", e2)
	}
}
