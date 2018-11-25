package command

type StopCommand struct {
	ap AudioPlayer
}

func NewStopCommand(ap AudioPlayer) Command {
	return &StopCommand{ap: ap}
}

func (sc *StopCommand) Execute() {
	sc.ap.Stop()
}
