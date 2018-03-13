package conf

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

func ReadFile(path string) (*Book, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	book := &Book{}
	book.Input = make([]Plugin, 0)
	book.Filter = make([]Plugin, 0)
	book.Output = make([]Plugin, 0)

	err = yaml.Unmarshal(raw, book)
	return book, err
}

type Plugin map[string]map[string]interface{}

type Book struct {
	Input  []Plugin `yaml:"input"`
	Filter []Plugin `yaml:"filter"`
	Output []Plugin `yaml:"output"`
}

func (p Plugin) Validate() error {
	if len(p) > 1 {
		return errors.New("Your object must be flat, with just one key")
	}
	return nil
}

func (p Plugin) Name() string {
	for k, _ := range p {
		return k
	}
	return ""
}

func (p Plugin) Args() map[string]interface{} {
	for _, v := range p {
		return v
	}
	return make(map[string]interface{})
}

type Configurable interface {
	Configure(conf map[string]interface{}) error
}

func ParseString(conf map[string]interface{}, key string, mandatory bool) (string, bool, error) {
	raw, ok := conf[key]
	if !ok {
		if mandatory {
			return "", false, fmt.Errorf("%s is mandatory in %s", key, conf)
		} else {
			return "", false, nil
		}
	}
	value, ok := raw.(string)
	if !ok {
		return "", false, fmt.Errorf("%s must be a string", key)
	}
	return value, true, nil
}

func ParseMapStringString(conf map[string]interface{}, key string, mandatory bool) (map[string]string, bool, error) {
	raw, ok := conf[key]
	if !ok {
		if mandatory {
			return nil, false, fmt.Errorf("%s is mandatory", key)
		} else {
			return nil, false, nil
		}
	}
	value, ok := raw.(map[string]string)
	if !ok {
		return nil, false, fmt.Errorf("%s must be a map of string string", key)
	}
	return value, true, nil
}

func ParseInt(conf map[string]interface{}, key string, mandatory bool) (int, bool, error) {
	raw, ok := conf[key]
	if !ok {
		if mandatory {
			return 0, false, fmt.Errorf("%s is mandatory", key)
		} else {
			return 0, false, nil
		}
	}
	value, ok := raw.(int)
	if !ok {
		return 0, false, fmt.Errorf("%s must be an integer", key)
	}
	return value, true, nil
}

func ParseTime(conf map[string]interface{}, key string, mandatory bool) (*time.Time, bool, error) {
	raw, ok := conf[key]
	if !ok {
		if mandatory {
			return nil, false, fmt.Errorf("%s is mandatory", key)
		} else {
			return nil, false, nil
		}
	}
	value, ok := raw.(*time.Time)
	if !ok {
		return nil, false, fmt.Errorf("%s must be a time.Time, did you parse it? : %s", key, raw)
	}
	return value, true, nil
}

func ParseArrayString(conf map[string]interface{}, key string, mandatory bool) ([]string, bool, error) {
	raw, ok := conf[key]
	if !ok {
		if mandatory {
			return nil, false, fmt.Errorf("%s is mandatory", key)
		} else {
			return nil, false, nil
		}
	}
	value, ok := raw.([]string)
	if !ok {
		return nil, false, fmt.Errorf("%s must be an array of strings: %s", key, raw)
	}
	return value, true, nil
}
