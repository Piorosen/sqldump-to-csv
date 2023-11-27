package sqldumptoobject

type Result struct {
	TableName string
	Columns   []string
	Values    [][]string
}
