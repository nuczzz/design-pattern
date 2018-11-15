package static

type Friend struct {
	Girl
	Myself
}

func (f *Friend)SendMsg2Girl(msg string) {
	f.Girl.ReadMsg(msg)
}

func (f *Friend)SendMsg2Boy(msg string){

}
