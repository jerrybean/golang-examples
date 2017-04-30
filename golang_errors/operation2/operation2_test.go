package operaion2

import (
	"testing"
)

func TestOperation2(t *testing.T) {
	opt := NewOperation2(1, 0)
	result, err := opt.Divide()
	if err != nil {
		t.Log(err)
	}
	t.Log(result)
}
