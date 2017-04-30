package operation3

import (
	"testing"
)

func TestOperation3(t *testing.T) {
	opt := NewOperation3(1, 0)
	result := opt.Divide()
	if opt.Err() != nil {
		t.Errorf(opt.Err().Error())
	}
	t.Log(result)
}
