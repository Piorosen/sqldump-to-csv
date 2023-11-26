package sqldumptocsv_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	sqldumptocsv "github.com/Piorosen/sqldump-to-csv"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func TestParse(t *testing.T) {
	// TODO
	read, err := os.Open("aaa.sql")
	if err != nil {
		t.Error(err, "not found aaa.sql")
	}
	defer read.Close()

	file, err := io.ReadAll(read)
	if err != nil {
		t.Error(err, "not readable aaa.sql")
	}

	fmt.Printf("%s", sqldumptocsv.Parse(file, 5))
}

func TestHomeParse(t *testing.T) {
	// TODO
	read, err := os.Open("chacha.sql")
	if err != nil {
		t.Error(err, "not found chacha.sql")
	}
	defer read.Close()

	file, err := io.ReadAll(read)
	if err != nil {
		t.Error(err, "not readable chacha.sql")
	}

	fmt.Printf("%s", sqldumptocsv.Parse(file, 5))
}

func TestLibr(t *testing.T) {
	data, err := sqlparser.Parse("SELECT * FROM TEST")
	if err != nil {
		t.Error(err, "not found chacha.sql")
	}
	fmt.Printf("stmt : %+v", data)
}
