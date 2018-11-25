package command

import "fmt"

type AudioPlayer struct {
}

func (ap *AudioPlayer) Play() {
	fmt.Println("play...")
}

func (ap *AudioPlayer) Rewind() {
	fmt.Println("rewind...")
}

func (ap *AudioPlayer) Stop() {
	fmt.Println("stop")
}
