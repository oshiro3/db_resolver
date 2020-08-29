package cli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/oshiro3/db_resolver/util"

	// blank import
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func createDiffFile(up, down, targetDb, wd string) string {
	version := time.Now().Format("06010215")
	f, err := os.Create(fmt.Sprintf("%s/%s/sqls/%s_migrate.up.sql", wd, targetDb, version))
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	fmt.Fprintln(f, up)
	f2, _ := os.Create(fmt.Sprintf("%s/%s/sqls/%s_migrate.down.sql", wd, targetDb, version))
	defer f2.Close()
	fmt.Fprintln(f2, down)
	return version
}

func createMigrationInstance(targetDb, wd string) (*migrate.Migrate, error) {
	user := util.GetOptionalEnvironmentVar("DATABASE_USERNAME", "docker")
	passwd := util.GetOptionalEnvironmentVar("DATABASE_PASSWD", "docker")
	host := util.GetOptionalEnvironmentVar("DATABASE_HOST", "db")
	m, err := migrate.New(
		fmt.Sprintf("file://%s", fmt.Sprintf("%s/%s/sqls", wd, targetDb)),
		fmt.Sprintf("mysql://%s:%s@tcp(%s:3306)/%s", user, passwd, host, targetDb))
	if err != nil {
		log.Println("err connection", err)
		return m, err
	}
	return m, nil
}

func up2Version(migrationInstance *migrate.Migrate, version string) error {
	iv, _ := strconv.Atoi(version)
	var uv = uint(iv)
	e := migrationInstance.Migrate(uv)
	if e != nil {
		log.Println("err migrate", e)
		return e
	}
	return nil
}

func up2Latest(migrationInstance *migrate.Migrate) error {
	e := migrationInstance.Up()
	if e != nil {
		log.Println("err migrate", e)
		return e
	}
	return nil
}

func setVersion(version, targetDb, wd string) {
	f, _ := os.Create(fmt.Sprintf("%s/%s/version", wd, targetDb))
	defer f.Close()
	fmt.Fprintf(f, version)
}

func setNewSchema(newSchema, targetDb, wd string) {
	e := os.Remove(fmt.Sprintf("%s/%s/out_of_date_schema.sql", wd, targetDb))
	if e != nil {
		fmt.Println(e)
	}
	os.Rename(
		fmt.Sprintf("%s/%s/schema.sql", wd, targetDb),
		fmt.Sprintf("%s/%s/out_of_date_schema.sql", wd, targetDb))
	f, err := os.Create(fmt.Sprintf("%s/%s/schema.sql", wd, targetDb))
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(f, newSchema)
	defer f.Close()
}
