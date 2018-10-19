package input

import (
	"github.com/factorysh/go-lepsius/model"
	//"gopkg.in/mcuadros/go-syslog.v2/format"
	"fmt"
	"gopkg.in/mcuadros/go-syslog.v2"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

func init() {
	register("syslog", &Syslog{})
}

type Syslog struct {
	protocol int
	host     string
	server   *syslog.Server
}

const (
	tcp = iota
	udp
	unix
)

type SyslogConf struct {
	Listen string
}

func (s *Syslog) Configure(c map[string]interface{}) error {
	var conf SyslogConf
	err := mapstructure.Decode(c, &conf)
	if err != nil {
		return err
	}
	u, err := url.Parse(conf.Listen)
	if err != nil {
		return err
	}
	switch u.Scheme {
	case "udp":
		s.protocol = udp
	case "tcp":
		s.protocol = tcp
	case "unix":
		s.protocol = unix
	default:
		return fmt.Errorf("Bad scheme : %s", u.Scheme)
	}
	s.host = u.Host
	s.server = syslog.NewServer()
	s.server.SetFormat(syslog.RFC5424)
	return nil
}

func (s *Syslog) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	var err error
	switch s.protocol {
	case tcp:
		err = s.server.ListenTCP(s.host)
	case udp:
		err = s.server.ListenUDP(s.host)
	case unix:
		err = s.server.ListenUnixgram(s.host)
	}
	if err != nil {
		panic(err)
	}
	s.server.Boot()
	return lines
}
