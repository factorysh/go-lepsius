package lepsius

import (
	"fmt"

	"github.com/vjeantet/grok"
	//"gopkg.in/mcuadros/go-syslog.v2"
	"gopkg.in/mcuadros/go-syslog.v2/format"
)

type SyslogGrokHandler struct {
	grok *grok.Grok
}

func NewHandler(pattern string) *SyslogGrokHandler {
	g, _ := grok.New()
	h := &SyslogGrokHandler{grok: g}
	h.grok.AddPattern("LEPSIUS", pattern)
	return h
}

func (s *SyslogGrokHandler) Handle(parts format.LogParts, machin int64, err error) {
	fmt.Println("parts", parts)
	s.grok.Parse("%{LEPSIUS}", "plop")
}
