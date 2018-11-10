package hook

type TemplateHook interface {
	BeforeDo()
	Do()
	AfterDo()
}

type Hook struct {
	SystemAction TemplateHook
	ClientAction TemplateHook
}

func (h *Hook) Do() {
	h.ClientAction.BeforeDo()
	h.SystemAction.Do()
	h.ClientAction.AfterDo()
}
