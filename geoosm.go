/**
 *Copyright (c) 2017 Nicolas Stuhlfauth
 *
 *This software is released under the MIT License.
 *https://opensource.org/licenses/MIT
 */

package geoosm

import (
	"encoding/json"
	"errors"
	"net/http"
)

// OSMData Stores information of OSM JSON response
type OSMData []struct {
	PlaceID     string   `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       string   `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
}

// OSMSlice Stores OSMData
type OSMSlice OSMData

// NewOSMData Generates a new OSMSlice
func NewOSMData() *OSMSlice {
	osm := make(OSMSlice, 0)
	return &osm
}

// GetJSON Writes JSON response to OSMData object
func (osmdata *OSMSlice) GetJSON(query string) (OSMData, error) {
	resp, err := http.Get("http://nominatim.openstreetmap.org/search?q=" + query + "&format=json")
	if err != nil {
		return nil, errors.New("Error: http.Get not possible")
	}
	err = json.NewDecoder(resp.Body).Decode(&osmdata)
	if err != nil {
		return nil, errors.New("Error: JSON Decoding not possible. Do you habe a Syntax error in your query?\nQuery URL: http://nominatim.openstreetmap.org/search?q=" + query + "&format=json")
	}
	data := OSMData(*osmdata)

	return data, nil
}
