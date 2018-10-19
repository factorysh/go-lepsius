package filter

import (
	"testing"
	"time"

	"github.com/factorysh/go-lepsius/model"
	"github.com/stretchr/testify/assert"
)

func TestHaproxy_date_format(t *testing.T) {
	dp := &DateParser{
		config: &DateParserConfig{
			Field:  "date",
			Layout: "Jan _2 15:04:05",
		},
	}
	line := model.Line{
		Values: map[string]interface{}{
			"date": "Oct  8 21:40:39",
		},
	}
	err := dp.Filter(&line)
	assert.Nil(t, err)
	tt, ok := line.Values["date"].(*time.Time)
	assert.True(t, ok)
	t.Log(tt)
}
