package tick

import "encoding/json"

type Line map[string]interface{}
type Parser func([]byte) (Line, error)

func JsonParser(raw []byte) (Line, error) {
	var o Line
	err := json.Unmarshal(raw, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func NoParser(raw []byte) (Line, error) {
	return Line{"message": raw}, nil
}
