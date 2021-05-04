package tokens

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Servicer {
	return &Service{db: db}
}

func (s *Service) Validate(header Header) error {
	return nil
}
