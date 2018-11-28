package adapter

import (
	"testing"
	"fmt"
)

func TestAdapter(t *testing.T) {
	videoPlay := NewVideoPlayer()
	moviePlayer := NewMoviePlayer()

	videoPlay.Play()
	fmt.Println()
	moviePlayer.Play()
}
