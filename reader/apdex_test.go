package reader

import (
	"testing"
	"time"
)

func TestApdex(t *testing.T) {
	a := &Apdex{}
	err := a.Configure(map[string]interface{}{
		"T":          500,
		"status_key": "status",
		"time_key":   "time",
		"date_key":   "timestamp",
		"tags":       []string{"host"},
	})
	if err != nil {
		t.Error(err)
	}
	go func() {
		evt := <-a.events
		t.Log("Event", evt)
	}()
	now := time.Now()
	err = a.Read(map[string]interface{}{
		"status":    "200",
		"timestamp": &now,
		"time":      342,
		"host":      "www.example.com",
	})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(1 * time.Second)
}
