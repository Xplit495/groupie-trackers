package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	MemberBis    string
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
	///Relations  	 []struct
}

func FetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {

		}
	}(resp.Body)

	var artists []Artist
	if err2 := json.NewDecoder(resp.Body).Decode(&artists); err2 != nil {
		return nil, err2
	}
	return artists, nil
}

func FetchArtistDetails(artistID int) (Artist, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", artistID)

	resp, err := http.Get(url)
	if err != nil {
		return Artist{}, err
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return Artist{}, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	var artist Artist
	if err1 := json.NewDecoder(resp.Body).Decode(&artist); err1 != nil {
		return Artist{}, err1
	}

	return artist, nil
}
