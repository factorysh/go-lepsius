package model

import (
	"errors"
	"fmt"
)

type Line struct {
	Data map[string]interface{}
}

func NewLine(datas ...interface{}) (*Line, error) {
	if len(datas)%2 != 0 {
		return nil, errors.New("Need an even number of arguments")
	}
	l := &Line{
		Data: make(map[string]interface{}),
	}

	for i := 0; i < len(datas)/2; i++ {
		k, ok := datas[i].(string)
		if !ok {
			return nil, fmt.Errorf("This key is not a string : %v", datas[i])
		}
		l.Data[k] = datas[i+1]
	}
	return l, nil
}
