# JSONResp

[![Godoc Reference][godoc-img]][godoc]
[![Build Status][ci-img]][ci]
[![Go Report Card][goreportcard-img]][goreportcard]

Package for generating JSON output from structs

## Usage
``` $ go get github.com/ewwwwwqm/jsonresp ```

### Example #1
``` main.go ```
```go
package main

import (
    "fmt"
    "github.com/ewwwwwqm/jsonresp"
)

func main() {
    obj := jsonresp.Output{"name": "Max"}
    fmt.Println(obj.String())
}
```

Output:
```json
{"name":"Max"}
```

### Example #2
``` main.go ```
```go
package main

import (
    "fmt"
    "net/http"
    "log"

    "github.com/ewwwwwqm/jsonresp"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	obj := jsonresp.Output{"message": "Hello there!", "number": 777}
    fmt.Fprintf(w, obj.String())
}

func main() {
    http.HandleFunc("/", index)
    err := http.ListenAndServe(":3001", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

``` $ go run main.go ```

Open http://127.0.0.1:3001/ via browser and check the output:
```json
{
    "message": "Hello there!",
    "number": 777
}
```

[godoc]: http://godoc.org/github.com/ewwwwwqm/jsonresp
[godoc-img]: https://godoc.org/github.com/ewwwwwqm/jsonresp?status.svg
[ci-img]: https://travis-ci.org/ewwwwwqm/jsonresp.svg?branch=master
[cov-img]: https://coveralls.io/repos/github/ewwwwwqm/jsonrespbadge.svg?branch=master
[ci]: https://travis-ci.org/ewwwwwqm/jsonresp
[cov]: https://coveralls.io/github/ewwwwwqm/jsonresp?branch=master
[goreportcard-img]: https://goreportcard.com/badge/github.com/ewwwwwqm/jsonresp
[goreportcard]: https://goreportcard.com/report/github.com/ewwwwwqm/jsonresp
