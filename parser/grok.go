package parser

import (
	_conf "github.com/bearstech/go-lepsius/conf"
	"github.com/vjeantet/grok"
)

type Grok struct {
	grok    *grok.Grok
	pattern string
}

func (g *Grok) Configure(conf map[string]interface{}) error {
	var err error
	g.grok, err = grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
	if err != nil {
		return err
	}

	err = g.grok.AddPatternsFromMap(map[string]string{
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

	path, ok, err := _conf.ParseString(conf, "path", false)
	if err != nil {
		return err
	}
	if ok {
		err = g.grok.AddPatternsFromPath(path)
		if err != nil {
			return err
		}
	}

	patterns, ok, err := _conf.ParseMapStringString(conf, "patterns", false)
	if err != nil {
		return err
	}
	if ok {
		for k, v := range patterns {
			err = g.grok.AddPattern(k, v)
			if err != nil {
				return err
			}
		}
	}
	g.pattern, _, err = _conf.ParseString(conf, "pattern", true)
	return err
}

func (g *Grok) Parse(line string) (map[string]string, error) {
	v, err := g.grok.Parse(g.pattern, line)
	if err != nil {
		return nil, err
	}
	return v, nil
}
