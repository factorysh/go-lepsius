package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {
	j := Filter["json"]
	err := j.Configure(map[string]interface{}{
		"field": "beuha",
	})
	assert.Nil(t, err)
	data := map[string]interface{}{
		"beuha": `{"age": 42}`,
	}
	err = j.Filter(data)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"beuha": map[string]interface{}{
			"age": 42.0,
		},
	}, data)

}
