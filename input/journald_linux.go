package input

import (
	"github.com/bearstech/go-lepsius/model"
	"github.com/coreos/go-systemd/sdjournal"
	"github.com/mitchellh/mapstructure"
	"time"
)

type Journald struct {
	journald *sdjournal.Journal
}

func (i *Journald) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	return lines
}

type JournaldConf struct {
	since   time.Duration
	Matches []sdjournal.Match
}

func (i *Journald) Configure(conf map[string]interface{}) error {
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
