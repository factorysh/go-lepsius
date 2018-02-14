package reader

import (
	"fmt"
)

type Stdout struct {
}

func (s *Stdout) Configure(conf map[string]interface{}) error {
	return nil
}

func (s *Stdout) Read(evt map[string]interface{}) error {
	fmt.Println(evt)
	return nil
}
