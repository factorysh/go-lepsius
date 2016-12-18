package lepsius

import (
	client "log/syslog"
	"os"
	"testing"
	"time"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func TestSyslog(t *testing.T) {
	_ = os.Remove("/tmp/test.sock")

	handler := NewHandler("%{HAPROXYHTTP}")

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	server.ListenUnixgram("/tmp/test.sock")
	server.Boot()

	go server.Wait()

	c, err := client.Dial("unixgram", "/tmp/test.sock", client.LOG_LOCAL0, "lepsius")
	if err != nil {
		t.Fatal(err)
	}
	c.Info(`Oct 29 23:59:33 my-server haproxy[16914]: 78.40.125.71:36602 [29/Oct/2015:23:59:29.957] http-in~ httpd/backend1 2488/0/0/1313/3801 200 423 - - ---- 1/1/0/1/0 0/0 "GET /test.php HTTP/1.1"`)
	time.Sleep(100 * time.Millisecond)

}
