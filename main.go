// Application which greets you.
package main

import (
	"fmt"
	"github.com/{{.repository.organization}}/{{.repository.name}}/internal/{{.repository.name}}"
)

func main() {
	fmt.Println(greet())
	{{.repository.name}}.PrintVersion()
}

func greet() string {
	return "Hi!"
}
