package output

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/mitchellh/mapstructure"
)

func init() {
	register("raven", &Raven{})
}

type Raven struct {
	client *raven.Client
	field  string
}

type RavenConf struct {
	Dsn          string
	FieldMessage string `mapstructure:"field_message"`
}

func (r *Raven) Configure(conf map[string]interface{}) error {
	var cfg RavenConf
	err := mapstructure.Decode(conf, cfg)
	if err != nil {
		return err
	}
	if cfg.FieldMessage == "" {
		cfg.FieldMessage = "message"
	}
	r.field = cfg.FieldMessage
	r.client, err = raven.New(cfg.Dsn)
	return err
}

type Interface struct {
	Context map[string]interface{}
}

func (i *Interface) Class() string {
	return "lepsius"
}

func (r *Raven) Read(evt map[string]interface{}) error {
	raw, ok := evt[r.field]
	if !ok {
		return fmt.Errorf("Can't find field %s", r.field)
	}
	msg, ok := raw.(string)
	if !ok {
		return fmt.Errorf("field %s is not a string : %v", r.field, raw)
	}
	r.client.CaptureMessage(msg, nil, &Interface{evt})
	return nil
}
