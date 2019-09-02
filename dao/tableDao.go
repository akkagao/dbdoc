package dao

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"dbdoc/models"
)

type tableDao struct {
}

/*
获取table列表
*/
func (d *tableDao) Tables(c *gin.Context) ([]string, error) {
	sqlShowTables := "SHOW TABLES"
	rows, err := DbEngin.Raw(sqlShowTables).Rows()
	if err != nil {
		return nil, err
	}

	tablesNames := []string{}
	tableName := ""
	for rows.Next() {
		rows.Scan(&tableName)
		tablesNames = append(tablesNames, tableName)
	}
	return tablesNames, nil
}

/*
获取表字典
*/
func (d *tableDao) TableDetail(c *gin.Context, tableName string) ([]models.TableField, error) {
	sqlTableDesc := "show full fields from %v"

	rows, err := DbEngin.Raw(fmt.Sprintf(sqlTableDesc, tableName)).Rows()
	if err != nil {
		return nil, err
	}

	tableFields := []models.TableField{}
	Field, Type, Collation, Null, Key, Default, Extra, Privileges, Comment := []byte{}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{}, []byte{}
	for rows.Next() {
		rows.Scan(&Field, &Type, &Collation, &Null, &Key, &Default, &Extra, &Privileges, &Comment)
		tableFields = append(tableFields, models.TableField{
			Field:      string(Field),
			Type:       string(Type),
			Collation:  string(Collation),
			Null:       string(Null),
			Key:        string(Key),
			Default:    string(Default),
			Extra:      string(Extra),
			Privileges: string(Privileges),
			Comment:    string(Comment),
		})
		// fmt.Println(string(Field), Type, Collation, Null, Key, Default, Extra, Privileges, Comment)
		// Field, Type, Collation, Null, Key, Default, Extra, Privileges, Comment = "", "", "", "", "", "", "", "", ""
	}
	return tableFields, nil
}
