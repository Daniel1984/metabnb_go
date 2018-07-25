package models

// Listing represents listings metadata from airbnb endpoint
type Listing struct {
	ExploreTabs []struct {
		Sections []struct {
			Listings []struct {
				Listing struct {
					BathroomLabel         string  `json:"bathroom_label"`
					Bathrooms             float32 `json:"bathrooms"`
					BedLabel              string  `json:"bed_label"`
					BedroomLabel          string  `json:"bedroom_label"`
					Bedrooms              float32 `json:"bedrooms"`
					Beds                  float32 `json:"beds"`
					City                  string  `json:"city"`
					ID                    int64   `json:"id"`
					IsNewListing          bool    `json:"is_new_listing"`
					IsSuperhost           bool    `json:"is_superhost"`
					Lat                   float64 `json:"lat"`
					Lng                   float64 `json:"lng"`
					LocalizedCity         string  `json:"localized_city"`
					LocalizedNeighborhood string  `json:"localized_neighborhood"`
					Name                  string  `json:"name"`
					Neighborhood          string  `json:"neighborhood"`
					PersonCapacity        int     `json:"person_capacity"`
					PictureCount          int     `json:"picture_count"`
					PictureURL            string  `json:"picture_url"`
					Picture               struct {
						LargeRo string `json:"large_ro"`
					} `json:"picture"`
				} `json:"listing"`
			} `json:"listings"`
		} `json:"sections"`
		PaginationMetadata struct {
			HasNextPage   bool `json:"has_next_page"`
			SectionOffset int  `json:"section_offset"`
		} `json:"pagination_metadata"`
	} `json:"explore_tabs"`
}
