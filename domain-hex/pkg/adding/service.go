package adding

import (
	"errors"
)

// ErrUnknown is used when a beer could not be found.
var ErrDuplicate = errors.New("beer already exists")

// Service provides beer adding operations.
type Service interface {
	AddBeer(...Beer)
	AddSampleBeers()
}

// Repository provides access to beer repository.
type Repository interface {
	// AddBeer saves a given beer to the repository.
	AddBeer(Beer) error
}

type service struct {
	bR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddBeer adds the given beer(s) to the database
func (s *service) AddBeer(b ...Beer) {

	// any validation can be done here

	for _, beer := range b {
		_ = s.bR.AddBeer(beer) // error handling omitted for simplicity
	}
}

// AddSampleBeers adds some sample beers to the database
func (s *service) AddSampleBeers() {

	// any validation can be done here

	for _, b := range DefaultBeers {
		_ = s.bR.AddBeer(b) // error handling omitted for simplicity
	}
}