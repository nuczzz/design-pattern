package command

type PlayCommand struct {
	ap AudioPlayer
}

func NewPlayCommand(ap AudioPlayer) Command {
	return &PlayCommand{ap: ap}
}

func (pc *PlayCommand) Execute() {
	pc.ap.Play()
}
