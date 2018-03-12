package parser

import (
	"encoding/json"
)

func init() {
	register("json", &Json{})
}

type Json struct {
}

func (j *Json) Configure(conf map[string]interface{}) error {
	return nil
}

func (j *Json) Parse(input []byte) (map[string]interface{}, error) {
	kv := make(map[string]interface{})
	err := json.Unmarshal(input, &kv)
	if err != nil {
		return nil, err
	}
	return kv, nil
}
