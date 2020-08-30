package cli

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		bytes, _ := ioutil.ReadFile("./shemas/new_schema.sql")
		u, d := cmp(string(bytes), "test_database", getWorkDir())
		assert.Equal(t, "\nBEGIN;\n\nSET FOREIGN_KEY_CHECKS = 0;\n\nDROP TABLE `test_table`;\n\nSET FOREIGN_KEY_CHECKS = 1;\n\nCOMMIT;", u)
		assert.Equal(t, "\nBEGIN;\n\nSET FOREIGN_KEY_CHECKS = 0;\n\nCREATE TABLE `test_table` (\n`id` INT (20) NOT NULL,\n`name` VARCHAR (20) COLLATE `utf8_bin` NOT NULL,\n`created_at` DATETIME DEFAULT NULL,\n`updated_at` DATETIME DEFAULT NULL\n) ENGINE = InnoDB, DEFAULT CHARACTER SET = utf8, DEFAULT COLLATE = utf8_bin;\n\nSET FOREIGN_KEY_CHECKS = 1;\n\nCOMMIT;", d)
	})
}
