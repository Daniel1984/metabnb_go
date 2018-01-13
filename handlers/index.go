package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// ListingsMetadata represents listings metadata from airbnb endpoint
type ListingsMetadata struct {
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

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return getErr
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	jsonErr := json.Unmarshal(body, target)

	if jsonErr != nil {
		return jsonErr
	}

	return nil
}

func getListingsURL(location string) string {
	return fmt.Sprintf("https://www.airbnb.com/api/v2/explore_tabs?version=1.3.2&_format=for_explore_search_web&experiences_per_grid=20&items_per_grid=50&guidebooks_per_grid=0&auto_ib=true&fetch_filters=true&is_guided_search=false&is_new_trips_cards_experiment=true&is_new_homes_cards_experiment=false&luxury_pre_launch=false&screen_size=large&show_groupings=false&supports_for_you_v3=true&timezone_offset=120&metadata_only=false&is_standard_search=true&selected_tab_id=all_tab&tab_id=home_tab&location=%v&federated_search_session_id=e30fad3d-4dfd-4348-b72a-bb2d1f53ca0c&_intents=p1&screen_size=large&key=d306zoyjsyarp7ifhu67rjxn52tv0t20&currency=USD&locale=en", location)
}

// Index root path handler
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	listingsMetadata := ListingsMetadata{}
	url := getListingsURL("london")

	err := getJSON(url, &listingsMetadata)
	if err != nil {
		fmt.Println(err)
	}

	json, errMarshal := json.Marshal(listingsMetadata)
	if errMarshal != nil {
		fmt.Println(err)
	}

	w.Write(json)
}
