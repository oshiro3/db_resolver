package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/schemalex/schemalex/diff"

	"io/ioutil"
	"log"
	"os"
)

// diff.Strings は ALTER を含むファイルを渡すと結果に nil を返すので注意

func compare() {
	new_schema, err := ioutil.ReadFile("./schemas/new_schema.sql")
	if err != nil {
		log.Printf("read err: %v\n", err)
	}
	old_schema, err2 := ioutil.ReadFile("./schemas/old_schema.sql")
	if err2 != nil {
		log.Printf("read err: %v\n", err2)
	}
	f, err3 := os.Create("sqls/02_diff_schema.up.sql")
	if err3 != nil {
		log.Printf("read err: %v\n", err3)
	}
	defer f.Close()
	if err5 := diff.Strings(f, string(old_schema), string(new_schema), diff.WithTransaction(true)); err != nil {
		log.Printf("diff err: %v\n", err5)
	}

	f2, err4 := os.Create("sqls/02_diff_schema.down.sql")
	if err4 != nil {
		log.Printf("read err: %v\n", err4)
	}
	defer f2.Close()
	diff.Strings(f2, string(new_schema), string(old_schema), diff.WithTransaction(true))
}
