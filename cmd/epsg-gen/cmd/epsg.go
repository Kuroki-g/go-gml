package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiBase = "https://apps.epsg.org/api/v1"

type crsEntry struct {
	Code int
	Name string
	Dim  uint
}

type coordRefSystemResp struct {
	Code     int       `json:"Code"`
	Name     string    `json:"Name"`
	CoordSys childLink `json:"CoordSys"`
}

type childLink struct {
	Code int `json:"Code"`
}

type coordSystemResp struct {
	Dimension int `json:"Dimension"`
}

// fetchEntry calls the EPSG API and returns the CRS dimension for the given EPSG code.
// Two requests are made: CoordRefSystem/{code} then CoordSystem/{CoordSys.Code}.
func fetchEntry(epsgCode int) (crsEntry, error) {
	var crs coordRefSystemResp
	if err := getJSON(fmt.Sprintf("%s/CoordRefSystem/%d", apiBase, epsgCode), &crs); err != nil {
		return crsEntry{}, err
	}
	var cs coordSystemResp
	if err := getJSON(fmt.Sprintf("%s/CoordSystem/%d", apiBase, crs.CoordSys.Code), &cs); err != nil {
		return crsEntry{}, err
	}
	return crsEntry{Code: epsgCode, Name: crs.Name, Dim: uint(cs.Dimension)}, nil
}

func getJSON(url string, v any) error {
	resp, err := http.Get(url) //nolint:noctx
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, url)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
