package main

type ProductRepository interface {
	GetStockCount(productID int) (int, error)
}

type AdService struct {
	repo ProductRepository
}

func NewAdService(r ProductRepository) *AdService {
	return &AdService{
		repo: r,
	}
}

func (s *AdService) CanShowAd(productID int) (bool, error) {
	count, err := s.repo.GetStockCount(productID)
	if err != nil {
		return false, err
	}

	return count > 10, nil
}