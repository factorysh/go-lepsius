package filter

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func init() {
	register("json", &Json{})
}

type Json struct {
	Field string
}

type JsonConfig struct {
	Field string
}

func (j *Json) Configure(conf map[string]interface{}) error {
	var cfg JsonConfig
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}
	j.Field = cfg.Field
	return nil
}

func (j *Json) Filter(line map[string]interface{}) error {
	raw, ok := line[j.Field]
	if ok {
		f, ok := raw.(string)
		if !ok {
			return fmt.Errorf("Only string can be parsed as JSON : %v", raw)
		}
		kv := make(map[string]interface{})
		err := json.Unmarshal([]byte(f), &kv)
		if err != nil {
			return err
		}
		line[j.Field] = kv
	}
	return nil
}
