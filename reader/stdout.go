package reader

import (
	"github.com/davecgh/go-spew/spew"
)

type Stdout struct {
}

func (s *Stdout) Configure(conf map[string]interface{}) error {
	return nil
}

func (s *Stdout) Read(evt map[string]interface{}) error {
	spew.Dump(evt)
	return nil
}
