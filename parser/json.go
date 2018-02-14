package parser

import (
	"encoding/json"
)

type Json struct {
}

func (j *Json) Configure(conf map[string]interface{}) error {
	return nil
}

func (j *Json) Parse(line string) (map[string]interface{}, error) {
	kv := make(map[string]interface{})
	err := json.Unmarshal([]byte(line), &kv)
	if err != nil {
		return nil, err
	}
	return kv, nil
}
