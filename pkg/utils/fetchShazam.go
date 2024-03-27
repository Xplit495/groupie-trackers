package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"unicode"
)

// ShazamResponse represents the structure of a Shazam API response.
type ShazamResponse struct {
	Tracks struct {
		Hits []struct {
			Track struct {
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Images   struct {
					Image string `json:"coverart"`
				} `json:"images"`
				Hub struct {
					Options []struct {
						Actions []struct {
							Uri string `json:"uri"`
						} `json:"actions"`
					} `json:"options"`
				} `json:"hub"`
			} `json:"track"`
			Snippet string `json:"snippet"`
		} `json:"hits"`
	} `json:"tracks"`
	Singers struct {
		Hits []struct {
			Singer struct {
				Name   string `json:"name"`
				Avatar string `json:"avatar"`
			} `json:"artist"`
		} `json:"hits"`
	} `json:"artists"`
}

// FetchShazam fetches Shazam data for a given input.
func FetchShazam(shazamInput string) []byte {
	// Clean the input to remove non-alphanumeric characters
	cleanShazamInput := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return ' '
	}, shazamInput)

	// Construct the URL for the Shazam API request
	url := fmt.Sprintf("https://shazam.p.rapidapi.com/search?term=%s&locale=fr-FR&offset=0&limit=6", cleanShazamInput)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Add necessary headers for the RapidAPI
	req.Header.Add("X-RapidAPI-Key", "f6c32c8c75msh764ee496b5cc231p1721f1jsn9e9dd8fb6363")
	req.Header.Add("X-RapidAPI-Host", "shazam.p.rapidapi.com")

	// Send the HTTP request
	res, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Fatal(err1)
		return nil
	}

	// Ensure the response body is closed after the function returns
	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			// Handle error if closing response body fails
		}
	}(res.Body)

	// Read the response body
	data, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		log.Fatal(err3)
		return nil
	}

	// Define a variable to hold the decoded JSON data
	var shazamData ShazamResponse

	// Unmarshal the JSON data into the ShazamResponse struct
	if err4 := json.Unmarshal(data, &shazamData); err4 != nil {
		log.Fatal(err4)
	}

	// Marshal the ShazamResponse struct back to JSON format
	jsonShazam, err5 := json.Marshal(shazamData)
	if err5 != nil {
		log.Fatal(err5)
	}

	return jsonShazam
}
