package filter

import (
	"testing"

	"github.com/factorysh/go-lepsius/model"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	j := Filter["json"]
	err := j.Configure(map[string]interface{}{
		"field": "beuha",
	})
	assert.Nil(t, err)
	data := model.Line{
		Values: map[string]interface{}{
			"beuha": `{"age": 42}`,
		},
	}
	err = j.Filter(&data)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"beuha": map[string]interface{}{
			"age": 42.0,
		},
	}, data.Values)

}
