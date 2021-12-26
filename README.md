[![Go](https://github.com/tabee/ahvvalidieren/actions/workflows/go.yml/badge.svg)](https://github.com/tabee/ahvvalidieren/actions/workflows/go.yml)
# swiss ahv nr validieren
validate a swiss ahv nr given in a string, returns true or false.
Checks the country code and ean13 checksum.

```go
package main

import c "github.com/tabee/ahvvalidieren"

func main() {
	arr := [2]string{"56.9217.0769.84", "756.3903.6825.80"}
	for j := 0; j < len(arr); j++ {
		v, _ := c.Validate(arr[j])
		println(v)
	}
}
```
You will see:
```bash
false
true
```