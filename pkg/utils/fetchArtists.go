package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Artist represents information about an artist.
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	MemberBis    string   `json:"memberBis"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// FetchArtists retrieves a list of artists from a remote API.
func FetchArtists() []Artist {
	// Fetch data from the remote API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	// Handle error if HTTP request fails
	if err != nil {
		fmt.Println("Error fetching artists: ", err)
		return nil
	}

	// Ensure response body is closed when function returns
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println("Failed to close response body:", err1)
			return
		}
	}(resp.Body)

	// Decode JSON response into a slice of Artist structs
	var artists []Artist
	if err2 := json.NewDecoder(resp.Body).Decode(&artists); err2 != nil {
		fmt.Println("Error parsing JSON: ", err2)
		return nil
	}

	// Concatenate members into a single string separated by comma and newline
	for i := 0; i < len(artists); i++ {
		artists[i].MemberBis = strings.Join(artists[i].Members, ",\n")
	}

	return artists
}
