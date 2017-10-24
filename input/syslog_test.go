package input

import (
	"testing"
)

func TestParse(t *testing.T) {
	cfg := make(map[string]interface{})
	cfg["listen"] = "udp://0.0.0.0:6379"
	s := &Syslog{}
	err := s.Configure(cfg)
	if err != nil {
		t.Error(err)
	}

}
