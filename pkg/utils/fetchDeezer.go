package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// FetchDeezer fetches Deezer artist data by artist name.
func FetchDeezer(artistName string) string {
	// URL encode the artist name to handle special characters properly
	encodedArtistName := url.QueryEscape(artistName)

	// Send GET request to Deezer API to fetch artist data
	resp, err := http.Get("https://api.deezer.com/search/artist?q=" + encodedArtistName)

	// Handle error if HTTP request fails
	if err != nil {
		fmt.Println("Error fetching artists: ", err)
		return ""
	}

	// Ensure response body is closed when function returns
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println("Failed to close response body:", err1)
			return
		}
	}(resp.Body)

	// Read the response body
	deezerData, _ := io.ReadAll(resp.Body)

	var deezerID string
	exit := false

	// Parse response body to extract Deezer artist ID
	for i := 0; i < len(deezerData); i++ {
		if exit == true {
			break
		}
		// Is here to extract the deezer artist ID from the Deezer API response
		if string(deezerData[i]) == "i" && string(deezerData[i+1]) == "d" {
			for j := i + 4; j < 100; j++ {
				if rune(deezerData[j]) >= '0' && rune(deezerData[j]) <= '9' {
					deezerID += string(deezerData[j])
				} else {
					exit = true
					break
				}
			}
		}
	}

	return deezerID
}
