package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/julienschmidt/httprouter"
	"metabnb/models"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getListingsURL(location string) string {
	return fmt.Sprintf("https://www.airbnb.com/api/v2/explore_tabs"+
		"?version=1.3.2"+
		"&_format=for_explore_search_web"+
		"&experiences_per_grid=20"+
		"&items_per_grid=50"+
		"&guidebooks_per_grid=0"+
		"&auto_ib=true"+
		"&fetch_filters=true"+
		"&is_guided_search=false"+
		"&is_new_trips_cards_experiment=true"+
		"&is_new_homes_cards_experiment=false"+
		"&luxury_pre_launch=false"+
		"&screen_size=large"+
		"&show_groupings=false"+
		"&supports_for_you_v3=true"+
		"&timezone_offset=120"+
		"&metadata_only=false"+
		"&is_standard_search=true"+
		"&selected_tab_id=all_tab"+
		"&tab_id=home_tab"+
		"&location=%v"+
		"&federated_search_session_id=e30fad3d-4dfd-4348-b72a-bb2d1f53ca0c"+
		"&_intents=p1"+
		"&screen_size=large"+
		"&key=d306zoyjsyarp7ifhu67rjxn52tv0t20"+
		"&currency=USD"+
		"&locale=en", location)
}

func getResponseBody(url string, ch chan []byte) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	ch <- body
}

// Listing root path handler
func GetListings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	queryValues := r.URL.Query()
	street := queryValues.Get("street")

	if street != "" {
		streetName := &url.URL{Path: street}
		encodedStreetName := streetName.String()
		url := getListingsURL(encodedStreetName)

		channel := make(chan []byte)
		go getResponseBody(url, channel)
		responseBody := <-channel

		listing := models.Listing{}
		json.Unmarshal(responseBody, &listing)
		fmt.Println(listing)

		json, errMarshal := json.Marshal(listing)
		if errMarshal != nil {
			fmt.Println(errMarshal)
			http.Error(w, http.StatusText(403), http.StatusForbidden)
			return
		}

		w.Write(json)
	} else {
		http.Error(w, http.StatusText(403), http.StatusForbidden)
	}
}
