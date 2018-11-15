package static

// Proxy Pattern Definition:
// Provide a surrogate or placeholder for another object to control access to it.

type Proxy interface {
	SendMsg2Girl(string)
	SendMsg2Boy(string)
}
