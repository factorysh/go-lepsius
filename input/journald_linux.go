package input

import (
	_conf "github.com/bearstech/go-lepsius/conf"
	"github.com/bearstech/go-lepsius/model"
	"github.com/coreos/go-systemd/sdjournal"
	"github.com/mitchellh/mapstructure"
)

type Journald struct {
	journald *sdjournal.Journal
}

func (i *Tail) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	return lines
}

type JournaldConf struct {
	since   time.Duration
	Matches []sdjournal.Match
}

func (i *Tail) Configure(conf map[string]interface{}) error {
	var jconf JournaldConf
	decoder, err := mapstructure.NewDecoder(mapstructure.DecoderConfig{
		Result: &jconf,
	})
	if err != nil {
		return err
	}
	err = decoder.Decode(conf)
	if err != nil {
		return err
	}
	return nil
}
