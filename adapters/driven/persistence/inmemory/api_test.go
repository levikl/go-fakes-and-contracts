package inmemory

import (
	"testing"

	"github.com/levikl/go-fakes-and-contracts/domain/planner"
)

func TestInmemoryAPI1(t *testing.T) {
	planner.API1Contract{NewAPI1: func() planner.API1 {
		return NewAPI1()
	}}.Test(t)
}
