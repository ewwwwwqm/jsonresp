package jsonresp_test

import (
	"fmt"

	"github.com/ewwwwwqm/jsonresp"
)

func ExampleOutput() {
	obj := jsonresp.Output{"name": "Max"}
	
	fmt.Println(obj.String())
	// Output:
	// {"name":"Max"}
}
