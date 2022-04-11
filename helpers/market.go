package helpers

import (
	"math"

	"github.com/steschwa/hopper-analytics-collector/models"
)

func ListingToListingDocument(listing models.Listing) models.ListingDocument {
	val, _ := listing.Price.Float64()

	return models.ListingDocument{
		Id:        listing.Id,
		Enabled:   listing.Enabled,
		Sold:      listing.Sold,
		Price:     val * math.Pow(10, -18),
		Buyer:     listing.Buyer,
		Seller:    listing.Seller,
		Timestamp: listing.Timestamp,
		HopperId:  listing.HopperId,
	}
}
