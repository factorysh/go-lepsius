package tick

import (
	"encoding/json"

	"gitlab.bearstech.com/bearstech/go-lepsius/tick/model"
)

type Parser func([]byte) (*model.Line, error)

func JsonParser(raw []byte) (*model.Line, error) {
	o := &model.Line{}
	err := json.Unmarshal(raw, o.Data)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func NoParser(raw []byte) (*model.Line, error) {
	return &model.Line{
		Data: map[string]interface{}{
			"message": raw},
	}, nil
}
