package command

type Keyboard struct {
	PlayCommand   Command
	RewindCommand Command
	StopCommand   Command
}

func (k *Keyboard) Play() {
	k.PlayCommand.Execute()
}

func (k *Keyboard) Rewind() {
	k.RewindCommand.Execute()
}

func (k *Keyboard) Stop() {
	k.StopCommand.Execute()
}
