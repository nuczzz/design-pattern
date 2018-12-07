package visitor

type AccountBook struct {
	List []Bill
}

func (a *AccountBook)AddBill(bill Bill) {
	a.List = append(a.List, bill)
}

func (a *AccountBook) Show(view AccountBookView) {
	for i := range a.List {
		a.List[i].Accept(view)
	}
}

func NewAccountBook() *AccountBook {
	return &AccountBook{
		List: make([]Bill, 0),
	}
}
