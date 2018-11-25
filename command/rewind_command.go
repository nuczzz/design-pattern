package command

type RewindCommand struct {
	ap AudioPlayer
}

func NewRewindCommand(ap AudioPlayer) Command {
	return &RewindCommand{ap: ap}
}

func (rc *RewindCommand) Execute() {
	rc.ap.Rewind()
}
