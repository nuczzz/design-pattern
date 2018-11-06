package simplefactory

import (
	"fmt"
	"time"
)

// Factory Pattern Definition:
// Define an interface for creating an object,
// but let subclasses decide which class to instantiate.
// Factory Method lets a class defer instantiation to subclasses.
type Node interface {
	GetRole() string
	GetIp() string
	GetStatus() string
}

// take and example of hadoop nameNode and dataNode,
// we can get different information with different node.
func NewNodeMachine(role, ip string) Node {
	switch role {
	case "nn":
		return &NameNode{
			role: "NameNode",
			ip:   ip,
		}
	case "dn":
		return &DataNode{
			role: "DataNode",
			ip:   ip,
		}
	default:
		panic(fmt.Sprintf("not support role: %v", role))
	}
}

// hadoop NameNode
type NameNode struct {
	ip   string
	role string
}

func (nn *NameNode) GetRole() string {
	return nn.role
}

func (nn *NameNode) GetIp() string {
	return nn.ip
}

func (nn *NameNode) GetStatus() string {
	if time.Now().Unix()%2 == 0 {
		return "I'm crash"
	}

	return "I'm running"
}

// hadoop DataNode
type DataNode struct {
	role string
	ip   string
}

func (dn *DataNode) GetRole() string {
	return dn.role
}

func (dn *DataNode) GetIp() string {
	return dn.ip
}

func (dn *DataNode) GetStatus() string {
	if time.Now().Unix()%2 == 0 {
		return "disk error: no space left on device"
	}

	return "healthy disk space"
}
