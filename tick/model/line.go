package model

import (
	"errors"
	"fmt"
	"time"
)

// Line of log
type Line struct {
	Data      map[string]interface{}
	KeepIt    bool
	Tags      map[string]string
	Timestamp time.Time
}

// NewLine from list of key, value, key, value…
func NewLine(datas ...interface{}) (*Line, error) {
	if len(datas)%2 != 0 {
		return nil, errors.New("Need an even number of arguments")
	}
	l := &Line{
		Data:   make(map[string]interface{}),
		KeepIt: true,
	}

	for i := 0; i < len(datas); i += 2 {
		k, ok := datas[i].(string)
		if !ok {
			return nil, fmt.Errorf("This key is not a string : %v", datas[i])
		}
		l.Data[k] = datas[i+1]
	}
	return l, nil
}

// Flatten Line as a map
func (l *Line) Flatten() map[string]interface{} {
	d := l.Data
	d["tags"] = l.Tags
	return d
}
