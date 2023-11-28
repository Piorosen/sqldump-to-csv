package sqldumptoobject

import (
	"fmt"

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

	var idx []int

	// set column is mapping data for insert.
	// but. not set value on column is insert data listed sequentially.
	// so. branch out for insert mode.
	if len(t.Columns) == 0 {
		idx = make([]int, len(table.Columns))
		for i := 0; i < len(idx); i++ {
			idx[i] = i
		}
	} else {
		idx = make([]int, len(table.Columns))
		for i := 0; i < len(idx); i++ {
			idx[i] = -1
		}

		for ti, tc := range table.Columns {
			for ci, c := range t.Columns {
				if c.String() == tc {
					idx[ti] = ci
				}
			}
			fmt.Println(t.Columns)
		}
	}

	rows := t.Rows.(sqlparser.Values)
	for _, r := range rows {
		cols := len(r)
		value := make([]string, len(table.Columns))
		for i := 0; i < cols; i++ {
			if idx[i] == -1 {
				return table, err
			}

			sqlval := r[i].(*sqlparser.SQLVal)
			value[idx[i]] = string(sqlval.Val)
		}
		table.Value = append(table.Value, value)
	}

	return table, nil
}
