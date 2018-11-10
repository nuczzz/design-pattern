package hook

type SystemHook interface {
	Do()
}

type ClientHook interface {
	BeforeDo()
	AfterDo()
}

type Hook struct {
	SystemAction SystemHook
	ClientAction ClientHook
}

func (h *Hook) Do() {
	h.ClientAction.BeforeDo()
	h.SystemAction.Do()
	h.ClientAction.AfterDo()
}
