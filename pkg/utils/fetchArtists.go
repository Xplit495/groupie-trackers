package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	MemberBis    string   `json:"memberBis"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func FetchArtists() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Println("Error fetching artists: ", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println("Failed to close response body:", err1)
			return
		}
	}(resp.Body)

	var artists []Artist

	if err2 := json.NewDecoder(resp.Body).Decode(&artists); err2 != nil {
		fmt.Println("Error parsing JSON: ", err2)
		return nil
	}

	for i := 0; i < len(artists); i++ {
		artists[i].MemberBis = strings.Join(artists[i].Members, ",\n")
	}

	return artists
}
