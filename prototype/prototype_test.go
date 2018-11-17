package prototype

import "testing"

func TestPrototype(t *testing.T) {
	boy1 := &Boy{
		Name: "boy1",
		Height: 160,
		Weight:58,
	}
	boy1.Print()

	boy2 := boy1.Clone()
	boy2.SetName("boy2")
	boy2.SetWeight(70)
	boy2.Print()

	boy3 := boy2.Clone()
	boy3.SetName("boy3")
	boy3.SetHeight(180)
	boy3.Print()
}
