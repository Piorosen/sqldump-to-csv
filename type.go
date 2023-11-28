package sqldumptoobject

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

type Table struct {
	TableName string
	Columns   []string
	Value     [][]string
}

func ConvertTable(query string) (*Table, error) {
	table, err := sqlparser.Parse(query)
	if err != nil {
		return nil, err
	}

	t := table.(*sqlparser.CreateTable)

	c := []string{}
	for _, v := range t.Columns {
		c = append(c, v.Name)
	}

	return &Table{
		TableName: t.NewName.Name.String(),
		Columns:   c,
	}, nil
}

func ConvertInsert(table *Table, query string) (*Table, error) {
	insert, err := sqlparser.Parse(query)
	if err != nil {
		return table, err
	}
	t := insert.(*sqlparser.Insert)

	// set column is mapping data for insert.
	// but. not set value on column is insert data listed sequentially.
	// so. branch out for insert mode.
	if len(t.Columns) == 0 {
		rows := t.Rows.(sqlparser.Values)
		for _, r := range rows {
			cols := len(r)
			value := []string{}
			for i := 0; i < cols; i++ {
				sqlval := r[i].(*sqlparser.SQLVal)
				value = append(value, string(sqlval.Val))
			}
			table.Value = append(table.Value, value)
		}
	} else {

	}

	return table, nil
	// data, err := sqlparser.Parse("INSERT INTO `member` VALUES (10,'test','test','test','/src/assets/noProfile.png'),(11,'test1','test1','test1','/src/assets/noProfile.png'),(12,'test2','test2','test2','/src/assets/noProfile.png'),(13,'test3','test3','test3','/src/assets/noProfile.png'),(14,'test4','test4','test4','/src/assets/noProfile.png'),(15,'guddnr0421@naver.com','이형욱','752MBNTQFZ','/src/assets/noProfile.png');")
}
