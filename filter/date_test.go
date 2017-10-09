package filter

import (
	"testing"
)

func TestHaproxy_date_format(t *testing.T) {
	dp := &DateParser{"date", "Jan _2 15:04:05"}
	conf := new(map[string]interface{})
	conf["date"] = "Oct  8 21:40:39"
	err := dp.Filter(conf)
	if err != nil {
		t.Error(err)
	}
	t.Log("Date format:", conf["date"])
}
