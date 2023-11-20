package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	sqldumptocsv "github.com/Piorosen/sqldump-to-csv"
)

func main() {
	file := flag.String("file", "", "Input SQL file.")
	export := flag.String("to", "stdout", "export to data. (stdout, json, csv)")
	flag.Parse()

	if file == nil || export == nil {
		fmt.Println(flag.ErrHelp.Error())
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Println(err.Error())
	}

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
	}

	if *export == "stdout" {
		fmt.Printf("%s", sqldumptocsv.Parse(b))
	}
}
