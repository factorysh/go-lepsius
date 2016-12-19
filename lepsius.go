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

func haproxyGrok(g *grok.Grok) error {
	err := g.AddPatternsFromMap(map[string]string{
		"HAPROXYCAPTUREDREQUESTHEADERS":  `%{DATA:request_header_host}\|%{DATA:request_header_x_forwarded_for}\|%{DATA:request_header_accept_language}\|%{DATA:request_header_referer}\|%{DATA:request_header_user_agent}`,
		"HAPROXYCAPTUREDRESPONSEHEADERS": `%{DATA:response_header_content_type}\|%{DATA:response_header_content_encoding}\|%{DATA:response_header_cache_control}\|%{DATA:response_header_last_modified}`,
		"HAPROXYTIME":                    `%{HOUR:haproxy_hour}:%{MINUTE:haproxy_minute}(?::%{SECOND:haproxy_second})`,
		"HAPROXYDATE":                    `%{MONTHDAY:haproxy_monthday}/%{MONTH:haproxy_month}/%{YEAR:haproxy_year}:%{HAPROXYTIME:haproxy_time}.%{INT:haproxy_milliseconds}`,
		"HAPROXYHTTP":                    `%{SYSLOGTIMESTAMP:syslog_timestamp} %{IPORHOST:syslog_server} %{SYSLOGPROG}: %{IP:client_ip}:%{INT:client_port} \[%{HAPROXYDATE:accept_date}\] %{NOTSPACE:frontend_name} %{NOTSPACE:backend_name}/%{NOTSPACE:server_name} %{INT:time_request}/%{INT:time_queue}/%{INT:time_backend_connect}/%{INT:time_backend_response}/%{NOTSPACE:time_duration} %{INT:http_status_code} %{NOTSPACE:bytes_read} %{DATA:captured_request_cookie} %{DATA:captured_response_cookie} %{NOTSPACE:termination_state} %{INT:actconn}/%{INT:feconn}/%{INT:beconn}/%{INT:srvconn}/%{NOTSPACE:retries} %{INT:srv_queue}/%{INT:backend_queue} (\{%{HAPROXYCAPTUREDREQUESTHEADERS}\})?( )?(\{%{HAPROXYCAPTUREDRESPONSEHEADERS}\})?( )?"(<BADREQ>|(%{WORD:http_verb} (%{URIPROTO:http_proto}://)?(?:%{USER:http_user}(?::[^@]*)?@)?(?:%{URIHOST:http_host})?(?:%{URIPATHPARAM:http_request})?( HTTP/%{NUMBER:http_version})?))?"`,
		"HAPROXYHTTPDIRECT":              `%{INT:stuff} \[%{HAPROXYDATE:accept_date}\] %{NOTSPACE:frontend_name} %{NOTSPACE:backend_name}/%{NOTSPACE:server_name} %{INT:time_request}/%{INT:time_queue}/%{INT:time_backend_connect}/%{INT:time_backend_response}/%{NOTSPACE:time_duration} %{INT:http_status_code} %{NOTSPACE:bytes_read} %{DATA:captured_request_cookie} %{DATA:captured_response_cookie} %{NOTSPACE:termination_state} %{INT:actconn}/%{INT:feconn}/%{INT:beconn}/%{INT:srvconn}/%{NOTSPACE:retries} %{INT:srv_queue}/%{INT:backend_queue} (\{%{HAPROXYCAPTUREDREQUESTHEADERS}\})?( )?(\{%{HAPROXYCAPTUREDRESPONSEHEADERS}\})?( )?"(<BADREQ>|(%{WORD:http_verb} (%{URIPROTO:http_proto}://)?(?:%{USER:http_user}(?::[^@]*)?@)?(?:%{URIHOST:http_host})?(?:%{URIPATHPARAM:http_request})?( HTTP/%{NUMBER:http_version})?))?"`,
	})
	if err != nil {
		return err
	}
	fmt.Println("plop")
	return nil
}

func NewHandler(pattern string) (*SyslogGrokHandler, error) {
	config := &grok.Config{NamedCapturesOnly: true}
	g, err := grok.NewWithConfig(config)
	if err != nil {
		return nil, err
	}
	err = haproxyGrok(g)
	if err != nil {
		return nil, err
	}
	err = g.AddPattern("LEPSIUS", pattern)
	if err != nil {
		return nil, err
	}
	return &SyslogGrokHandler{grok: g}, nil
}

func (s *SyslogGrokHandler) Handle(parts format.LogParts, machin int64, err error) {
	for k, v := range parts {
		fmt.Println("## ", k, ": ", v)
	}
	//fmt.Println("grok", s.grok)
	values, err := s.grok.Parse("%{LEPSIUS}", parts["content"].(string))
	if err != nil {
		panic(err)
	}
	fmt.Println("grok values: ", values)
}
