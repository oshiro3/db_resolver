package cli

import (
	"bytes"
	"fmt"

	"github.com/oshiro3/db_resolver/find"
	"github.com/schemalex/schemalex/diff"
)

// diff.Strings は ALTER を含むファイルを渡すと結果に nil を返すので注意
func cmp(newSchema, targetDb, wd string) (string, string) {
	opts := find.Option{Src: "filesystem"}
	currentSchema := find.Find(opts, fmt.Sprintf("%s/%s/schema.sql", wd, targetDb))
	var upBuf bytes.Buffer
	var downBuf bytes.Buffer
	diff.Strings(&upBuf, currentSchema, newSchema, diff.WithTransaction(true))
	diff.Strings(&downBuf, newSchema, currentSchema, diff.WithTransaction(true))
	return upBuf.String(), downBuf.String()
}
