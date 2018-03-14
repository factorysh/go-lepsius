package filter

import (
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
	Field   string
	Pattern string
}

func (y *Yangtze) Configure(conf map[string]interface{}) error {
	var cfg YangtzeConf
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}
	y.config = &cfg
	return nil
}

func (y *Yangtze) Filter(line *model.Line) error {
	return nil
}
