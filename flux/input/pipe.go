package input

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"syscall"

	"github.com/containerd/fifo"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/semantic"
	"github.com/influxdata/flux/values"
)

const PipeKind = "pipe"

type PipeOpSpec struct {
	Path string `json:"path"`
}

func init() {
	pipeSignature := semantic.FunctionPolySignature{
		Parameters: map[string]semantic.PolyType{
			"path": semantic.String,
		},
		Required: semantic.LabelSet{"path"},
		Return:   flux.TableObjectType,
	}
	flux.RegisterPackageValue("input", PipeKind,
		flux.FunctionValue(PipeKind, createPipeOpSpec, pipeSignature))
	flux.RegisterOpSpec(PipeKind, newPipeOp)
	plan.RegisterProcedureSpec(PipeKind, newPipeProcedure, PipeKind)
}

func createPipeOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	spec := new(PipeOpSpec)
	if p, _, e := args.GetString("path"); e != nil {
		return nil, e
	} else {
		spec.Path = p
	}
	return spec, nil
}

func newPipeOp() flux.OperationSpec {
	return new(PipeOpSpec)
}

func (s *PipeOpSpec) Kind() flux.OperationKind {
	return PipeKind
}

type PipeProcedureSpec struct {
	plan.DefaultCost
	Path string
}

func newPipeProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	spec, ok := qs.(*PipeOpSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", qs)
	}

	return &PipeProcedureSpec{
		Path: spec.Path,
	}, nil
}

func (s *PipeProcedureSpec) Kind() plan.ProcedureKind {
	return PipeKind
}

func (s *PipeProcedureSpec) Copy() plan.ProcedureSpec {
	ns := new(PipeProcedureSpec)
	*ns = *s
	return ns
}

type PipeSource struct {
	alloc   *memory.Allocator
	fifoCtx context.Context
	fifo    io.ReadWriteCloser
	reader  *bufio.Reader
	Path    string
}

func createPipeSource(prSpec plan.ProcedureSpec, dsid execute.DatasetID,
	a execute.Administration) (execute.Source, error) {
	spec, ok := prSpec.(*PipeProcedureSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", prSpec)
	}

	p := NewPipeSource(a.Allocator())
	p.Path = spec.Path

	return execute.CreateSourceFromDecoder(p, dsid, a)
}

func NewPipeSource(a *memory.Allocator) *PipeSource {
	return &PipeSource{alloc: a}
}

func (p *PipeSource) Connect() error {
	p.fifoCtx = context.Background()
	var err error
	p.fifo, err = fifo.OpenFifo(p.fifoCtx, p.Path, syscall.O_CREAT+syscall.O_RDWR,
		0660)
	if err != nil {
		return err
	}
	p.reader = bufio.NewReader(p.fifo)
	return nil
}

func (p *PipeSource) Fetch() (bool, error) {
	return true, nil
}

func (p *PipeSource) Decode() (flux.Table, error) {
	line, err := p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	key := execute.NewGroupKey([]flux.ColMeta{}, []values.Value{})
	tb := execute.NewColListTableBuilder(key, p.alloc)
	tb.AppendString(0, line)
	return tb.Table()
}

func (p *PipeSource) Close() error {
	p.fifoCtx.Done()
	return p.fifo.Close()
}
