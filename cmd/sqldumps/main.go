package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	sqldumptocsv "github.com/Piorosen/sqldump-to-object"
)

func main() {
	file := flag.String("file", "", "Input SQL file.")
	export := flag.String("to", "stdout", "export to data. (stdout, json)")
	top := flag.Int("top", -1, "top K")

	flag.Parse()

	if file == nil || export == nil || top == nil {
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
	r := sqldumptocsv.Parse(b, *top)
	if *export == "stdout" {
		fmt.Printf("%s\n", r)
	} else if *export == "json" {
		j, err := json.Marshal(r)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("%s\n", string(j))
		}
	}
}
