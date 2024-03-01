package utils

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func FetchArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	if err1 := json.NewDecoder(resp.Body).Decode(&artists); err1 != nil {
		return nil, err1
	}

	return artists, nil
}
