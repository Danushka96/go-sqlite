package Migrations

import (
	"github.com/danushka96/go-sqlite/SQL"
	"github.com/danushka96/go-sqlite/SQL/Const"
	"github.com/danushka96/go-sqlite/SQL/Models"
)

func Seed() {
	table := Models.New("table1")
	tableField1 := Models.TableField{
		Name:       "id",
		DataType:   Const.Varchar,
		NotNull:    true,
		PrimaryKey: true,
	}
	tableField2 := Models.TableField{
		Name:       "name",
		DataType:   Const.Varchar,
		NotNull:    true,
		PrimaryKey: false,
	}
	table.SetTableFields(tableField1)
	table.SetTableFields(tableField2)
	SQL.CreateTable(table)
	m := make(map[string]string)
	m["id"] = "123"
	m["name"] = "name123"
	SQL.InsertData(table, m)
}
