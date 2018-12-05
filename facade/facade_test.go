package facade

import "testing"

func TestFacade(t *testing.T) {
	movie := NewMovie()
	movie.Play()
}
