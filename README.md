# Mariadb SQL Dump file to Object in memory.

## Intoduce

## How Use?

```go
// Install
$ go install github.com/Piorosen/sqldump-to-csv/cmd/sqldumps@v1.1.1
$ sqldumps -file test.sql -to json
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
