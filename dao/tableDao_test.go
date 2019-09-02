package dao

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_tableDao_Tables(t *testing.T) {
	c := InitTest()
	tables, err := TableDao.Tables(c)
	spew.Dump(err)
	spew.Dump(tables)
}

func Test_tableDao_TableDetail(t *testing.T) {
	c := InitTest()
	tableFields, err := TableDao.TableDetail(c, "user")
	spew.Dump(err)
	spew.Dump(tableFields)
}
