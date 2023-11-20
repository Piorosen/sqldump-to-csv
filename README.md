# Mariadb SQL Dump file to Object in memory.

## Intoduce

## How Use?

```go
// Install
$ go get github.com/Piorosen/sqldump-to-csv@v1.0.0
```

```go
package main

import sqldumptocsv "github.com/Piorosen/sqldump-to-csv"

func main() { 
    file, _ := os.Open("test.sql")
    obj, _ := io.ReadAll(file)
    r := sqldumptocsv.Parse(obj)
    fmt.Printf("%s", r)
}
```

## License 

MIT License
