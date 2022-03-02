// Application which greets you.
package main

import (
	"fmt"
	"github.com/{{.organization}}/{{.projectNmae}}/internal/{{.projectNmae}}"
)

func main() {
	fmt.Println(greet())
	{{.repository.name}}.PrintVersion()
}

func greet() string {
	return "Hi!"
}
