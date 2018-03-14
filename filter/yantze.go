package filter

import (
	"fmt"
	"github.com/athoune/yangtze/index"
	"github.com/bearstech/go-lepsius/model"
	"github.com/mitchellh/mapstructure"
)

func init() {
	register("yangtze", &Yangtze{})
}

type Yangtze struct {
	config *YangtzeConf
	index  *index.Index
}

type YangtzeConf struct {
	Field    string
	Patterns []string
	Target   string
}

func (y *Yangtze) Configure(conf map[string]interface{}) error {
	var cfg YangtzeConf
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}
	y.config = &cfg
	y.index, err = index.NewSimple()
	if err != nil {
		return err
	}
	for _, p := range cfg.Patterns {
		pp, err := y.index.Parser().Parse([]byte(p))
		if err != nil {
			return err
		}
		y.index.AddPattern(pp)
	}
	return nil
}

func (y *Yangtze) Filter(line *model.Line) error {
	raw, ok := line.Values[y.config.Field]
	if ok {
		f, ok := raw.(string)
		if !ok {
			return fmt.Errorf("Yangzte only handles string : %v", raw)
		}
		_, ok = y.index.ReadLine([]byte(f))
		line.Keep = ok
	}
	return nil
}
