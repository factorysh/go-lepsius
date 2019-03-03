package output

import (
	"fmt"

	"github.com/influxdata/flux/values"

	"github.com/influxdata/flux/execute"

	"github.com/davecgh/go-spew/spew"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/semantic"
)

const SpewKind = "spew"

func init() {
	spewSignature := flux.FunctionSignature(
		map[string]semantic.PolyType{
			"tables": semantic.Object,
		},
		nil,
	)
	flux.RegisterPackageValue("output", SpewKind,
		flux.FunctionValueWithSideEffect(SpewKind, createSpewOpSpec, spewSignature))
	flux.RegisterOpSpec(SpewKind, func() flux.OperationSpec { return &SpewOpSpec{} })
	plan.RegisterProcedureSpecWithSideEffect(SpewKind, newSpewProcedure, SpewKind)
	//execute.RegisterTransformation(SpewKind, createSpewTransformation)
}

func createSpewOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	spew.Dump("createSpewOpSpec", args)
	fmt.Println(args.GetAll())
	tables, _, err := args.GetObject("tables")
	if err != nil {
		return nil, err
	}
	spec := new(SpewOpSpec)
	spec.tables = tables
	return spec, nil
}

func newSpewOp() flux.OperationSpec {
	return new(SpewOpSpec)
}

type SpewOpSpec struct {
	tables values.Object
}

func (s *SpewOpSpec) Kind() flux.OperationKind {
	return SpewKind
}

type SpewProcedureSpec struct {
	plan.DefaultCost
	tables values.Object
}

func (s *SpewProcedureSpec) Kind() plan.ProcedureKind {
	return SpewKind
}

func (s *SpewProcedureSpec) Copy() plan.ProcedureSpec {
	ns := new(SpewProcedureSpec)
	*ns = *s
	return ns
}

func newSpewProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	spec, ok := qs.(*SpewOpSpec)
	fmt.Println(spec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", qs)
	}

	s := &SpewProcedureSpec{}
	s.tables = spec.tables
	return s, nil
}

func createSpewTransformation(id execute.DatasetID, mode execute.AccumulationMode, spec plan.ProcedureSpec, a execute.Administration) (execute.Transformation, execute.Dataset, error) {
	s, ok := spec.(*SpewProcedureSpec)
	if !ok {
		return nil, nil, fmt.Errorf("invalid spec type %T", spec)
	}
	cache := execute.NewTableBuilderCache(a.Allocator())
	d := execute.NewDataset(id, mode, cache)
	t := NewSpewTransformation(d, cache, s)
	return t, d, nil
}

type SpewTransformation struct {
	d     execute.Dataset
	cache execute.TableBuilderCache
	spec  *SpewProcedureSpec
}

func NewSpewTransformation(d execute.Dataset, cache execute.TableBuilderCache, spec *SpewProcedureSpec) *SpewTransformation {
	return &SpewTransformation{
		d:     d,
		cache: cache,
		spec:  spec,
	}
}

func (s *SpewTransformation) Finish(id execute.DatasetID, err error) {
	s.d.Finish(err)
}

func (s *SpewTransformation) Process(id execute.DatasetID, tbl flux.Table) error {
	spew.Dump("Process", tbl)
	return nil
}

func (s *SpewTransformation) RetractTable(id execute.DatasetID, key flux.GroupKey) error {
	return s.d.RetractTable(key)
}

func (s *SpewTransformation) UpdateProcessingTime(id execute.DatasetID, pt execute.Time) error {
	return s.d.UpdateProcessingTime(pt)
}

func (s *SpewTransformation) UpdateWatermark(id execute.DatasetID, pt execute.Time) error {
	return s.d.UpdateWatermark(pt)
}

/*
type SpewSource struct {
	alloc *memory.Allocator
}

func createSpewSource(prSpec plan.ProcedureSpec, dsid execute.DatasetID,
	a execute.Administration) (execute.Source, error) {
	spec, ok := prSpec.(*SpewProcedureSpec)
	fmt.Println(spec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", prSpec)
	}

	p := NewSpewSource(a.Allocator())

	return execute.CreateSourceFromDecoder(p, dsid, a)
}

func NewSpewSource(a *memory.Allocator) *SpewSource {
	return &SpewSource{alloc: a}
}

func (s *SpewSource) Connect() error {
	return nil
}

func (s *SpewSource) Close() error {
	return nil
}

func (s *SpewSource) Decode() (flux.Table, error) {
	return nil, nil
}

func (s *SpewSource) Fetch() (bool, error) {
	return true, nil
}
*/
