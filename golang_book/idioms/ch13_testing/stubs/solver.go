package stubs

import (
	"context"
	"errors"
	"io"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(ctx context.Context, expr string) (float64, error)
}

func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
	curExpr, err := readToNewLine(r)
	if err != nil {
		return 0, err
	}
	if len(curExpr) == 0 {
		return 0, errors.New("no expr to read")
	}
	ans, err := p.Solver.Resolve(ctx, curExpr)
	return ans, err
}

func readToNewLine(r io.Reader) (string, error) {
	var out []byte
	buf := make([]byte, 1)
	for {
		_, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				return string(buf), nil
			}
		}
		if buf[0] == '\n' {
			break
		}
		out = append(out, buf[0])
	}
	return string(out), nil
}

type MathSolverStub struct {}

func (ms MathSolverStub) Resolve(_ context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errors.New("invalid expression: ( 2 + 2 * 10")
	}
	return 0, nil	
}