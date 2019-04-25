package testing

import (
	"context"
	"os"
	"syscall"
	"testing"

	"github.com/influxdata/flux/lang"

	"github.com/containerd/fifo"
	_ "github.com/factorysh/go-lepsius/flux/input"
	"github.com/factorysh/go-lepsius/flux/query"
	_ "github.com/influxdata/flux/builtin"
	"github.com/stretchr/testify/assert"
)

func TestFifo(t *testing.T) {
	ctxFifo, cancel := context.WithCancel(context.Background())
	f, err := fifo.OpenFifo(ctxFifo, "/tmp/lepsius", syscall.O_CREAT+syscall.O_RDWR,
		0660)
	assert.NoError(t, err)
	f.Write([]byte("Hello world"))
	ql := `
	import "input"
	input.fifo(path:"/tmp/lepsius")
	`

	c := lang.FluxCompiler{Query: ql}

	querier, err := query.New(os.Stdout)
	assert.NoError(t, err)
	err = querier.Query(context.Background(), c)
	assert.NoError(t, err)

	f.Write([]byte("plop\n"))
	cancel()
}
