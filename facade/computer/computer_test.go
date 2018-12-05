package computer

import (
	"testing"
)

func TestComputer(t *testing.T) {
	computer1 := NewComputer("cpu1", "disk1", "memory1")
	computer2 := NewComputer("cpu2", "disk2", "memory2")

	computer1.Start()
	computer2.Start()

	computer2.Stop()
	computer1.Stop()
}
