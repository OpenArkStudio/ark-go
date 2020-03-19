package base_test

import (
	"github.com/ArkNX/ark-go/base"
	"testing"
)

func TestAFVector3D(t *testing.T) {
	v, err := base.NewAFVector3DFromString("21.23,3,5")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", v)
}
