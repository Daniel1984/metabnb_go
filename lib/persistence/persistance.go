package persistence

import "metabnb/models"

type DatabaseHandler interface {
	AddListing(models.Listing) ([]byte, error)
	FindListing([]byte) (models.Listing, error)
	FindListingById(string) (models.Listing, error)
	FindAllListings() ([]models.Listing, error)
}
