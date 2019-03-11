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

const FifoKind = "fifo"

type FifoOpSpec struct {
	Path string `json:"path"`
}

func init() {
	fifoSignature := semantic.FunctionPolySignature{
		Parameters: map[string]semantic.PolyType{
			"path": semantic.String,
		},
		Required: semantic.LabelSet{"path"},
		Return:   flux.TableObjectType,
	}
	flux.RegisterPackageValue("input", FifoKind,
		flux.FunctionValue(FifoKind, createFifoOpSpec, fifoSignature))
	flux.RegisterOpSpec(FifoKind, newFifoOp)
	plan.RegisterProcedureSpec(FifoKind, newFifoProcedure, FifoKind)
	execute.RegisterSource(FifoKind, createFifoSource)
}

func createFifoOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	spec := new(FifoOpSpec)
	p, _, e := args.GetString("path")
	if e != nil {
		return nil, e
	}
	spec.Path = p
	fmt.Println("createOpSpec")
	return spec, nil
}

func newFifoOp() flux.OperationSpec {
	return new(FifoOpSpec)
}

func (s *FifoOpSpec) Kind() flux.OperationKind {
	return FifoKind
}

type FifoProcedureSpec struct {
	plan.DefaultCost
	Path string
}

func newFifoProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	spec, ok := qs.(*FifoOpSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", qs)
	}

	fmt.Println("newFifoProcedure")
	return &FifoProcedureSpec{
		Path: spec.Path,
	}, nil
}

func (s *FifoProcedureSpec) Kind() plan.ProcedureKind {
	return FifoKind
}

func (s *FifoProcedureSpec) Copy() plan.ProcedureSpec {
	ns := new(FifoProcedureSpec)
	*ns = *s
	return ns
}

type FifoSource struct {
	alloc   *memory.Allocator
	fifoCtx context.Context
	fifo    io.ReadWriteCloser
	reader  *bufio.Reader
	line    []byte
	Path    string
}

func createFifoSource(prSpec plan.ProcedureSpec, dsid execute.DatasetID,
	a execute.Administration) (execute.Source, error) {
	spec, ok := prSpec.(*FifoProcedureSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", prSpec)
	}

	p := NewFifoSource(a.Allocator())
	p.Path = spec.Path

	fmt.Println("createFifoSource")
	return execute.CreateSourceFromDecoder(p, dsid, a)
}

func NewFifoSource(a *memory.Allocator) *FifoSource {
	return &FifoSource{alloc: a}
}

func (p *FifoSource) Connect() error {
	fmt.Println("Connect")
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

func (p *FifoSource) Fetch() (bool, error) {
	fmt.Println("Fetch")
	line, err := p.reader.ReadBytes('\n')
	if err != nil {
		return false, err
	}
	p.line = line
	return true, nil
}

func (p *FifoSource) Decode() (flux.Table, error) {
	fmt.Println("Decode")
	key := execute.NewGroupKey([]flux.ColMeta{}, []values.Value{})
	tb := execute.NewColListTableBuilder(key, p.alloc)
	tb.AppendString(0, string(p.line))
	return tb.Table()
}

func (p *FifoSource) Close() error {
	p.fifoCtx.Done()
	return p.fifo.Close()
}
