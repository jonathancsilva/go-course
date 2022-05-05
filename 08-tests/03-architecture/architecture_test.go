package architecture

import (
	"runtime"
	"testing"
)

func TestDependency(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Not work in architecture 64")
	}
	// ....
	t.Fail()
}
