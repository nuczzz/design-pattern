package template_method

import "testing"

func TestTemplateRun(t *testing.T) {
	var animal Template
	animal.Animal = NewElephant()
	animal.Run()

	animal.Animal = NewAnt()
	animal.Run()
}
