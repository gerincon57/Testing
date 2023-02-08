package products

// constructor
func NewStorageMock() *storageMock {
	return &storageMock{}
}

type storageMock struct {
	Data []Product
	Err  error
	Spy  bool
}

func (sm *storageMock) GetAllBySeller(sellerID string) (m []Product, err error) {
	sm.Spy = true
	m = sm.Data
	err = sm.Err
	return m, err
}
