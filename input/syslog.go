package input

import (
	_conf "github.com/bearstech/go-lepsius/conf"
	//"github.com/bearstech/go-lepsius/model"
	//"gopkg.in/mcuadros/go-syslog.v2/format"
	"fmt"
	"net/url"
)

type Syslog struct {
	protocol int
	host     string
}

const (
	tcp = iota
	udp
	unix
)

func (s *Syslog) Configure(conf map[string]interface{}) error {
	listen, _, err := _conf.ParseString(conf, "listen", true)
	u, err := url.Parse(listen)
	if err != nil {
		return err
	}
	fmt.Println(u)
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
	return nil
}

//func (s *Syslog) Lines() chan *model.Line {

//}
