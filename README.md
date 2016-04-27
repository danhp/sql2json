# sql2json

> Dump an sql table to JSON

## Install

    go get github.com/danhp/sql2json

## Usage

```go
func Run(rows *sql.Rows) (string, error)
```
 --
```go
package main

import(
    "github.com/phamdaniel/sql2json"
)

func main() {
    db := sql.Open(driverName, dataSourceName)
    rows := db.Query(query)

    jsonData, err := sql2json.Run(rows)
    if err != nil {
        // Handle it.
    }
}
```

Sample Output:
```json
[{"date":"2014-01-01T00:00:00Z","name":"Some Event","region":"Some place","signup":"http://challonge.com/","tid":1}]
```

## License
MIT © [Daniel Pham](https://danhp.github.io)
