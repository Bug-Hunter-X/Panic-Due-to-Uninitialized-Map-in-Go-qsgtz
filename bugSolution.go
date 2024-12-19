```go
package main

import "fmt"

func main() {
    var m = make(map[string]int) // Initialize the map
    m["key"] = 10
    fmt.Println(m["key"])
}
```