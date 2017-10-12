package reader

// https://en.wikipedia.org/wiki/Apdex

import (
	"fmt"
	_conf "github.com/bearstech/go-lepsius/conf"
	"time"
)

type Apdex struct {
	t          int
	status_key string
	time_key   string
	date_key   string
	events     chan *Event
	tags       []string
}

type Event struct {
	TimeStamp    *time.Time
	Satisfied    uint32
	Tolerating   uint32
	NotSatisfied uint32
	Tags         map[string]string
}

func (a *Apdex) Configure(conf map[string]interface{}) error {
	var err error
	a.t, _, err = _conf.ParseInt(conf, "T", true)
	if err != nil {
		return err
	}
	a.status_key, _, err = _conf.ParseString(conf, "status_key", true)
	if err != nil {
		return err
	}
	a.time_key, _, err = _conf.ParseString(conf, "time_key", true)
	if err != nil {
		return err
	}
	a.events = make(chan *Event)
	a.date_key, _, err = _conf.ParseString(conf, "date_key", true)
	if err != nil {
		return err
	}
	a.tags, _, err = _conf.ParseArrayString(conf, "tags", false)
	if err != nil {
		return err
	}
	return nil
}

func (a *Apdex) Read(evt map[string]interface{}) error {
	timestamp, _, err := _conf.ParseTime(evt, a.date_key, true)
	if err != nil {
		return err
	}
	status, _, err := _conf.ParseString(evt, a.status_key, true)
	if err != nil {
		return err
	}
	if len(status) != 3 {
		return fmt.Errorf("Bad status lenght : %s", status)
	}
	event := &Event{
		TimeStamp: timestamp,
		Tags:      make(map[string]string),
	}
	switch s := status[0]; s {
	case '1', '3', '4':
		// this request is not interesting, nor its speed
		return nil
	case '5':
		event.NotSatisfied += 1
	case '2':
		t, _, err := _conf.ParseInt(evt, a.time_key, true)
		if err != nil {
			return err
		}
		if t <= a.t {
			event.Satisfied += 1
		} else {
			if t <= a.t*4 {
				event.Tolerating += 1
			} else {
				event.NotSatisfied += 1
			}
		}
	default:
		return fmt.Errorf("Strange status : %s", status)
	}
	for _, tag := range a.tags {
		event.Tags[tag], _, err = _conf.ParseString(evt, tag, false)
		if err != nil {
			return err
		}
	}
	a.events <- event
	return nil
}
