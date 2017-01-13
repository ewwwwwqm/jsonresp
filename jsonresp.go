package jsonresp

import(
	"encoding/json"
)

type Output map[string]interface{}

func (r Output) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}