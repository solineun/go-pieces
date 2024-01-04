package stubs

import (
	"context"
	"strings"
	"testing"
)

func TestProcessorProcessExpression(t *testing.T) {
	p := Processor{MathSolverStub{}}
	in := strings.NewReader(
`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10`)
	data := []float64{22, 40, 0, 0}
	for _, d := range data {
		res, err := p.ProcessExpression(context.Background(), in)
		if err != nil {
			t.Error(err)
			continue
		}
		if res != d {
			t.Errorf("Expected result %f, got %f", d, res)
		}
	}
}

