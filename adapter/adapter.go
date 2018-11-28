package adapter

import "fmt"

// Adapter Pattern Definition:
// Convert the interface of a class into another interface clients expect.
// Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.

type Player interface {
	Play()
}

type VideoPlayer struct {
}

func NewVideoPlayer() Player {
	return &VideoPlayer{}
}

func (vp *VideoPlayer) Play() {
	fmt.Println("video play...")
}

type Movie struct {
}

func (m *Movie)  PlaySound() {
	fmt.Println("movie play sound...")
}

func (m *Movie) PlayPicture() {
	fmt.Println("movie play picture...")
}

type MovieAdapter struct {
	movie *Movie
}

func (ma *MovieAdapter) Play() {
	ma.movie.PlaySound()
	ma.movie.PlayPicture()
}

func NewMoviePlayer() Player {
	return &MovieAdapter{
		movie: &Movie{},
	}
}
