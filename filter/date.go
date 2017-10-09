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
	conf[d.field], err = time.Parse(d.layout, raw)
	return err
}
