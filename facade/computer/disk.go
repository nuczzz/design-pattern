package computer

import "fmt"

type disk struct {
	sn string
}

func (d *disk) start() {
	fmt.Printf("[%v] disk start...\n", d.sn)
}

func (d *disk) stop() {
	fmt.Printf("[%v] disk stop...\n", d.sn)
}

func newDisk(sn string) *disk {
	return &disk{sn: sn}
}
