package bridge

import "testing"

func TestBridge(t *testing.T) {
	bridge := NewBridgeInstance()
	bridge.SetInstance(NewInstanceA())
	bridge.Operation()

	bridge.SetInstance(NewInstanceB())
	bridge.Operation()
}
