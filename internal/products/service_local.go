package products

import "log"

type serviceLocal struct {
	repo Repository
}

func NewServiceLocal(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *serviceLocal) GetAllBySeller(sellerID string) ([]Product, error) {
	data, err := s.repo.GetAllBySeller(sellerID)
	if err != nil {
		log.Println("error in repository", err.Error(), "sellerId:", sellerID)
		return nil, err
	}
	return data, err
}
