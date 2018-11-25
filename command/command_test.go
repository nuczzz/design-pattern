package command

import "testing"

func TestCommand(t *testing.T) {
	audioPlayer := AudioPlayer{}

	playCommand := NewPlayCommand(audioPlayer)
	rewindCommand := NewRewindCommand(audioPlayer)
	stopCommand := NewStopCommand(audioPlayer)

	keyboard := Keyboard{
		PlayCommand:   playCommand,
		RewindCommand: rewindCommand,
		StopCommand:   stopCommand,
	}

	keyboard.Play()
	keyboard.Rewind()
	keyboard.Stop()
	keyboard.Play()
	keyboard.Stop()
}
