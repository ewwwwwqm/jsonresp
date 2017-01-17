// Package for generating JSON output from structs.
package jsonresp

import(
	"encoding/json"
)

// Output represents an interface of map[string]
type Output map[string]interface{}

// Marshal JSON and output as string
func (r Output) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}