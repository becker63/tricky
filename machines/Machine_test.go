package machine

import (
	"testing"

	memoryfs "testing/fstest"

	demo "github.com/becker63/tricky/Boxes/Demo"
)

func TestDemo(t *testing.T) {
	Demo := demo.Demo{}

	files := make(memoryfs.MapFS)

	NewMachine(files, Demo, "an id")
}
