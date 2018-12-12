package bridge

type Bridge interface {
	SetInstance(instance Instance)
	GetInstance() Instance
	Operation()
}

type InstanceBridge struct {
	instance Instance
}

func (i *InstanceBridge) SetInstance(instance Instance) {
	i.instance = instance
}

func (i *InstanceBridge) GetInstance() Instance {
	return i.instance
}

func (i *InstanceBridge) Operation() {
	i.instance.Operation()
}

func NewBridgeInstance() Bridge {
	return &InstanceBridge{}
}
