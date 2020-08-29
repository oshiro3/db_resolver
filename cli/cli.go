package cli

import (
	"fmt"
	"log"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/oshiro3/db_resolver/find"
)

// Up makes DB version latest
func Up(targetDb, src, path string) {
	wd := getWorkDir()
	opts := find.Option{Src: src}
	newSchema := find.Find(opts, path)
	up, down := cmp(newSchema, targetDb, wd)
	version := createDiffFile(up, down, targetDb, wd)
	m, _ := createMigrationInstance(targetDb, wd)
	err := up2Latest(m)
	if err != nil {
		log.Fatalf("fatal: %s", err)
	}
	setVersion(version, targetDb, wd) // rewrite version file
	setNewSchema(newSchema, targetDb, wd)
	log.Println(version)
}

// Set create directory tree under .../dbs
func Set(db string) {
	wd := getWorkDir()
	fmt.Println(wd)
	if err := os.MkdirAll(path.Join(wd, "dbs", db, "sqls"), 644); err != nil {
		log.Printf("err: %s", err)
	}
}

func getWorkDir() string {
	wd := os.Getenv("DB_RESOLVER_ROOT")
	if wd == "" {
		home, _ := homedir.Dir()
		wd = home
	}
	return path.Join(wd, "dbs")
}
