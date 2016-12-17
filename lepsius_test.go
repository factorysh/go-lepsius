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

	handler := NewHandler("")

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenUnixgram("/tmp/test.sock")
	server.Boot()

	go server.Wait()

	c, err := client.Dial("unixgram", "/tmp/test.sock", client.LOG_LOCAL0, "lepsius")
	if err != nil {
		t.Fatal(err)
	}
	c.Debug("Plop")
	time.Sleep(100 * time.Millisecond)

}
