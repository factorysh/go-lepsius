package parser

import (
	"fmt"

	_conf "github.com/factorysh/go-lepsius/conf"
	"github.com/vjeantet/grok"
)

func init() {
	register("grok", &Grok{})
}

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
		"USERNAME":       `[a-zA-Z0-9._-]+`,
		"USER":           `%{USERNAME}`,
		"EMAILLOCALPART": `[a-zA-Z][a-zA-Z0-9_.+-=:]+`,
		"EMAILADDRESS":   `%{EMAILLOCALPART}@%{HOSTNAME}`,
		"INT":            `(?:[+-]?(?:[0-9]+))`,

		"YEAR":   `(\d\d){1,2}`,
		"HOUR":   `(?:2[0123]|[01]?[0-9])`,
		"MINUTE": `(?:[0-5][0-9])`,
		// '60' is a leap second in most time standards and thus is valid.
		"SECOND":                         `(?:(?:[0-5]?[0-9]|60)(?:[:.,][0-9]+)?)`,
		"TIME":                           `%{HOUR}:%{MINUTE}(?::%{SECOND})([0-9])?`,
		"HTTPDATE":                       `%{MONTHDAY}/%{MONTH}/%{YEAR}:%{TIME} %{INT}`,
		"IPV6":                           `((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?`,
		"IPV4":                           `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`,
		"IP":                             `(?:%{IPV6}|%{IPV4})`,
		"HAPROXYCAPTUREDREQUESTHEADERS":  `%{DATA:request_header_host}\|%{DATA:request_header_x_forwarded_for}\|%{DATA:request_header_accept_language}\|%{DATA:request_header_referer}\|%{DATA:request_header_user_agent}`,
		"HAPROXYCAPTUREDRESPONSEHEADERS": `%{DATA:response_header_content_type}\|%{DATA:response_header_content_encoding}\|%{DATA:response_header_cache_control}\|%{DATA:response_header_last_modified}`,
		"HAPROXYTIME":                    `%{HOUR:haproxy_hour}:%{MINUTE:haproxy_minute}(?::%{SECOND:haproxy_second})`,
		"HAPROXYDATE":                    `%{MONTHDAY:haproxy_monthday}/%{MONTH:haproxy_month}/%{YEAR:haproxy_year}:%{HAPROXYTIME:haproxy_time}.%{INT:haproxy_milliseconds}`,
		"HAPROXYHTTP":                    `%{SYSLOGTIMESTAMP:syslog_timestamp} %{IPORHOST:syslog_server} %{SYSLOGPROG}: %{IP:client_ip}:%{INT:client_port} \[%{HAPROXYDATE:accept_date}\] %{NOTSPACE:frontend_name} %{NOTSPACE:backend_name}/%{NOTSPACE:server_name} %{INT:time_request}/%{INT:time_queue}/%{INT:time_backend_connect}/%{INT:time_backend_response}/%{NOTSPACE:time_duration} %{INT:http_status_code} %{NOTSPACE:bytes_read} %{DATA:captured_request_cookie} %{DATA:captured_response_cookie} %{NOTSPACE:termination_state} %{INT:actconn}/%{INT:feconn}/%{INT:beconn}/%{INT:srvconn}/%{NOTSPACE:retries} %{INT:srv_queue}/%{INT:backend_queue} (\{%{HAPROXYCAPTUREDREQUESTHEADERS}\})?( )?(\{%{HAPROXYCAPTUREDRESPONSEHEADERS}\})?( )?"(<BADREQ>|(%{WORD:http_verb} (%{URIPROTO:http_proto}://)?(?:%{USER:http_user}(?::[^@]*)?@)?(?:%{URIHOST:http_host})?(?:%{URIPATHPARAM:http_request})?( HTTP/%{NUMBER:http_version})?))?"`,
		"HAPROXYHTTPDIRECT":              `%{INT:stuff} \[%{HAPROXYDATE:accept_date}\] %{NOTSPACE:frontend_name} %{NOTSPACE:backend_name}/%{NOTSPACE:server_name} %{INT:time_request}/%{INT:time_queue}/%{INT:time_backend_connect}/%{INT:time_backend_response}/%{NOTSPACE:time_duration} %{INT:http_status_code} %{NOTSPACE:bytes_read} %{DATA:captured_request_cookie} %{DATA:captured_response_cookie} %{NOTSPACE:termination_state} %{INT:actconn}/%{INT:feconn}/%{INT:beconn}/%{INT:srvconn}/%{NOTSPACE:retries} %{INT:srv_queue}/%{INT:backend_queue} (\{%{HAPROXYCAPTUREDREQUESTHEADERS}\})?( )?(\{%{HAPROXYCAPTUREDRESPONSEHEADERS}\})?( )?"(<BADREQ>|(%{WORD:http_verb} (%{URIPROTO:http_proto}://)?(?:%{USER:http_user}(?::[^@]*)?@)?(?:%{URIHOST:http_host})?(?:%{URIPATHPARAM:http_request})?( HTTP/%{NUMBER:http_version})?))?"`,
		"CHRONO":                         `%{NUMBER}\ws`,
		"TRAEFIKCLF":                     `%{IP:client_ip} %{WORD:ident}|- %{WORD:auth}|- \[%{HTTPDATE:timestamp}\] "%{WORD:verb} %{NOTSPACE:request} HTTP/%{NUMBER:httpversion}" %{NUMBER:response} %{NUMBER:bytes} %{QS:referrer} %{QS:agent} %{NUMBER:request_count} %{QS:backend_container} %{QS:backend_address} %{CHRONO:chrono}`,
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

func (g *Grok) Parse(line []byte) (map[string]interface{}, error) {
	v, err := g.grok.Parse(g.pattern, string(line))
	if err != nil {
		return nil, err
	}
	fmt.Println("grok", v)
	vv := make(map[string]interface{})
	for key, value := range v {
		vv[key] = value
	}
	return vv, nil
}
