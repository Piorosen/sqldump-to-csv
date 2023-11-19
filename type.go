package sqldumptocsv

type Result struct {
	TableName string
	Columns   []string
	Values    [][]string
}
