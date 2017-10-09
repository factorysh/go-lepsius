package filter

import (
	_conf "github.com/bearstech/go-lepsius/conf"
	"testing"
)

func TestHaproxy_date_format(t *testing.T) {
	dp := &DateParser{"date", "Jan _2 15:04:05"}
	conf := make(map[string]interface{})
	conf["date"] = "Oct  8 21:40:39"
	err := dp.Filter(conf)
	if err != nil {
		t.Error(err)
	}
	t.Log("Date format:", conf["date"])
	tt, _, err := _conf.ParseTime(conf, "date", true)
	if err != nil {
		t.Error(err)
	}
	t.Log(tt)
}
