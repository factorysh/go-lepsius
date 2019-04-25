package query

import (
	"context"
	"io"
	"math"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/control"
	"github.com/influxdata/flux/csv"
)

type Querier struct {
	C      *control.Controller
	writer io.Writer
}

func New(writer io.Writer) (*Querier, error) {
	cfg := control.Config{
		ConcurrencyQuota:         1,
		MemoryBytesQuotaPerQuery: math.MaxInt64,
		QueueSize:                1,
	}
	ctrl, err := control.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Querier{
		C:      ctrl,
		writer: writer,
	}, nil

}

// Query does query
func (q *Querier) Query(ctx context.Context, c flux.Compiler) error {
	query, err := q.C.Query(ctx, c)
	if err != nil {
		return err
	}
	//fmt.Println(query.Err())
	results := flux.NewResultIteratorFromQuery(query)
	defer results.Release()
	//fmt.Println(results.Err())

	encoder := csv.NewMultiResultEncoder(csv.DefaultEncoderConfig())
	_, err = encoder.Encode(q.writer, results)
	return err
}
