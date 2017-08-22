package main

import (
	"encoding/json"
)

type Request struct {
	Method   string
	URL      string
	Body     string
	Encoding string

	Headers map[string]string
}

func ParseUA(s string) (r Request, err error) {
	err = json.Unmarshal([]byte(s), &r)

	return
}
