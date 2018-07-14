package pattern

type PatternInstance interface {
	//get the id of instance, different instance has different id
	GetInstanceId() int64
}
