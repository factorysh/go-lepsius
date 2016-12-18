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
	config := &grok.Config{NamedCapturesOnly: true}
	g, err := grok.NewWithConfig(config)
	if err != nil {
		panic(err)
	}
	h := &SyslogGrokHandler{grok: g}
	h.grok.AddPattern("HAPROXYTIME", `(?!<[0-9])%{HOUR:haproxy_hour}:%{MINUTE:haproxy_minute}(?::%{SECOND:haproxy_second})(?![0-9])`)
	h.grok.AddPattern("HAPROXYDATE", `%{MONTHDAY:haproxy_monthday}/%{MONTH:haproxy_month}/%{YEAR:haproxy_year}:%{HAPROXYTIME:haproxy_time}.%{INT:haproxy_milliseconds}`)
	h.grok.AddPattern("HAPROXYHTTP", `%{SYSLOGTIMESTAMP:syslog_timestamp} %{IPORHOST:syslog_server} %{SYSLOGPROG}: %{IP:client_ip}:%{INT:client_port} \[%{HAPROXYDATE:accept_date}\] %{NOTSPACE:frontend_name} %{NOTSPACE:backend_name}/%{NOTSPACE:server_name} %{INT:time_request}/%{INT:time_queue}/%{INT:time_backend_connect}/%{INT:time_backend_response}/%{NOTSPACE:time_duration} %{INT:http_status_code} %{NOTSPACE:bytes_read} %{DATA:captured_request_cookie} %{DATA:captured_response_cookie} %{NOTSPACE:termination_state} %{INT:actconn}/%{INT:feconn}/%{INT:beconn}/%{INT:srvconn}/%{NOTSPACE:retries} %{INT:srv_queue}/%{INT:backend_queue} (\{%{HAPROXYCAPTUREDREQUESTHEADERS}\})?( )?(\{%{HAPROXYCAPTUREDRESPONSEHEADERS}\})?( )?"(<BADREQ>|(%{WORD:http_verb} (%{URIPROTO:http_proto}://)?(?:%{USER:http_user}(?::[^@]*)?@)?(?:%{URIHOST:http_host})?(?:%{URIPATHPARAM:http_request})?( HTTP/%{NUMBER:http_version})?))?"`)

	h.grok.AddPattern("LEPSIUS", pattern)
	return h
}

func (s *SyslogGrokHandler) Handle(parts format.LogParts, machin int64, err error) {
	fmt.Println("parts", parts)
	//fmt.Println("grok", s.grok)
	values, err := s.grok.Parse("LEPSIUS", parts["content"].(string))
	if err != nil {
		panic(err)
	}
	fmt.Println("values", values)
}
