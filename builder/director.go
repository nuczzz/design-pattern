package builder

type HumanDirector struct {
}

func (d *HumanDirector) CreateHuman(hb HumanBuilder) Human {
	hb.BuildHead()
	hb.BuildBody()
	hb.BuildHand()
	hb.BuildFoot()
	return hb.NewHuman()
}

func NewHumanDirector() *HumanDirector {
	return &HumanDirector{}
}
