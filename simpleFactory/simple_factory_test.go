package simpleFactory

import (
	"testing"
)

func TestNew(t *testing.T) {
	New("car").Drive()
	New("ship").Drive()
	New("aircraft").Drive()
}
