package filter

import (
	_conf "github.com/bearstech/go-lepsius/conf"
	"github.com/mitchellh/mapstructure"
	"time"
)

func init() {
	register("dateparser", &DateParser{})
}

type DateParser struct {
	config *DateParserConfig
}

type DateParserConfig struct {
	field  string
	layout string
}

func (d *DateParser) Configure(conf map[string]interface{}) error {
	var c DateParserConfig
	return mapstructure.Decode(conf, &c)
}

func (d *DateParser) Filter(line map[string]interface{}) error {
	raw, _, err := _conf.ParseString(line, d.config.field, true)
	if err != nil {
		return err
	}
	t, err := time.Parse(d.config.layout, raw)
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
	line[d.config.field] = &t
	return err
}
