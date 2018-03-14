package filter

import (
	"fmt"
	"github.com/bearstech/go-lepsius/model"
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
	Field  string
	Layout string
}

func (d *DateParser) Configure(conf map[string]interface{}) error {
	var c DateParserConfig
	return mapstructure.Decode(conf, &c)
}

func (d *DateParser) Filter(line *model.Line) error {
	raw, ok := line.Values[d.config.Field]
	if ok {
		f, ok := raw.(string)
		if !ok {
			return fmt.Errorf("Only string can be parsed as date: %v", raw)
		}
		t, err := time.Parse(d.config.Layout, f)
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
		line.Values[d.config.Field] = &t
	}
	return nil
}
