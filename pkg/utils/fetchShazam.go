package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ArtistHit struct {
	Artist Singer `json:"artist"`
}

type ArtistsResponse struct {
	Hits []ArtistHit `json:"hits"`
}

type Singer struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Track struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    string `json:"image"`
	URL      string `json:"url"`
}

type TrackHit struct {
	Track Track `json:"track"`
}

type TracksResponse struct {
	Hits []TrackHit `json:"hits"`
}

type ShazamResponse struct {
	Singers ArtistsResponse `json:"artists"`
	Tracks  TracksResponse  `json:"tracks"`
}

func FetchShazam(shazamInput string) []byte {
	fmt.Println(shazamInput)

	urlLink := fmt.Sprintf("https://shazam.p.rapidapi.com/search?term=%s&locale=fr-FR&offset=0&limit=5", shazamInput)

	req, err := http.NewRequest("GET", urlLink, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-RapidAPI-Key", "f6c32c8c75msh764ee496b5cc231p1721f1jsn9e9dd8fb6363")
	req.Header.Add("X-RapidAPI-Host", "shazam.p.rapidapi.com")

	client := &http.Client{}

	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {

		}
	}(resp.Body)

	data, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println(string(data))

	var shazamData ShazamResponse

	if err4 := json.Unmarshal(data, &shazamData); err4 != nil {
		log.Fatal(err3)
	}

	jsonShazam, err5 := json.Marshal(shazamData)
	if err5 != nil {
		log.Fatal(err5)
	}

	return jsonShazam
}
