package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	NewElement("john", "red")
	NewElement("julia", "green")
	NewElement("john", "blue")
}
