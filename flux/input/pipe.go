package input

import (
	"fmt"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/semantic"
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
