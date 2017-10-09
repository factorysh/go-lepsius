package filter

import (
	_conf "github.com/bearstech/go-lepsius/conf"
	"time"
)

type DateParser struct {
	field  string
	layout string
}

func (d *DateParser) Configure(conf map[string]interface{}) error {
	var err error
	d.field, _, err = _conf.ParseString(conf, "field", false)
	if err != nil {
		return err
	}
	d.layout, _, err = _conf.ParseString(conf, "layout", false)
	if err != nil {
		return err
	}
	return nil
}

func (d *DateParser) Filter(conf map[string]interface{}) error {
	raw, _, err := _conf.ParseString(conf, d.field, true)
	if err != nil {
		return err
	}
	t, err := time.Parse(d.layout, raw)
	if err != nil {
		return err
	}
	if t.Year() == 0 {
		n := time.Now()
		if n.Month() > t.Month() {
			t = t.AddDate(n.Year()-1, 0, 0)
		} else {
			t = t.AddDate(n.Year(), 0, 0)
		}
	}
	conf[d.field] = t
	return err
}
