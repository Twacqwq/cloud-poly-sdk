package cloudpolysdk

type PolyCore struct {
	Stores *Store
}

func New(option ...Option) (*PolyCore, error) {
	cloudPoly := &PolyCore{
		Stores: &Store{},
	}

	for _, opt := range option {
		opt(cloudPoly.Stores)
		if cloudPoly.Stores.Err != nil {
			return nil, cloudPoly.Stores.Err
		}
	}

	return cloudPoly, nil
}
