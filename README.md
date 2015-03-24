# sql2json

> Dump an sql table to JSON

## Install

    go get github.com/phamdaniel/sql2json

## Usage

```go
func Run(rows *sql.Rows) ([]byte, error)
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

    stringData := string(jsonData)
}
```

Sample Output:
```json
[{"date":"2014-01-01T00:00:00Z","name":"Some Event","region":"Some place","signup":"http://challonge.com/","tid":1}]
```

## License
MIT © [Daniel Pham](https://phamdaniel.github.io)
