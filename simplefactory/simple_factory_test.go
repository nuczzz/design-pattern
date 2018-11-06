package simplefactory

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimpleFactory(t *testing.T) {
	nameNode := NewNodeMachine("nn", "192.168.0.100")
	dataNode1 := NewNodeMachine("dn", "192.168.0.200")
	dataNode2 := NewNodeMachine("dn", "192.168.0.201")

	for _, node := range []Node{nameNode, dataNode1, dataNode2} {
		time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)
		printNodeInfo(t, node)
	}
}

func printNodeInfo(t *testing.T, node Node) {
	t.Logf(`
		timestamp: %v
		role: %v
		ip: %v
		status: %v`, time.Now().Unix(), node.GetRole(), node.GetIp(), node.GetStatus())
}
