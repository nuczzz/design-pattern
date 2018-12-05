package facade

import "fmt"

// Facade Pattern Definition:
// Provide a unified interface to a set of interfaces in a subsystem.
// Facade defines a higher-level interface that makes the subsystem easier to use.

type Facade struct {
	M Music
	V Video
	C Comment
}

func (f *Facade) Play() {
	f.M.GetMusic()
	f.V.GetVideo()
	f.C.GetComments()
}

type Music struct {
}

func (m *Music) GetMusic() {
	fmt.Println("play music...")
}

type Video struct {
}

func (v *Video) GetVideo() {
	fmt.Println("play video...")
}

type Comment struct {
}

func (c *Comment) GetComments() {
	fmt.Println("show comments...")
}

func NewMovie() *Facade {
	return &Facade{}
}
