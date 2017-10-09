package reader

// https://en.wikipedia.org/wiki/Apdex

import (
	_conf "github.com/bearstech/go-lepsius/conf"
	"time"
)

type Apdex struct {
	t          int
	status_key string
	time_key   string
	date_key   time.Time
}

func (a *Apdex) Configure(conf map[string]interface{}) error {
	var err error
	a.t, _, err = _conf.ParseInt(conf, "T", true)
	if err != nil {
		return err
	}
	a.status_key, _, err = _conf.ParseString(conf, "status_key", false)
	if err != nil {
		return err
	}
	a.time_key, _, err = _conf.ParseString(conf, "time_key", false)
	if err != nil {
		return err
	}
	return nil
}

func (a *Apdex) Read(map[string]string) error {

	return nil
}
