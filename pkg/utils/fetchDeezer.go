package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func FetchDeezer(artistName string) string {

	encodedArtistName := url.QueryEscape(artistName)

	resp, err := http.Get("https://api.deezer.com/search/artist?q=" + encodedArtistName)

	if err != nil {
		fmt.Println("Error fetching artists: ", err)
		return ""
	}

	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println("Failed to close response body:", err1)
			return
		}
	}(resp.Body)

	deezerData, _ := io.ReadAll(resp.Body)

	var deezerID string
	exit := false

	for i := 0; i < len(deezerData); i++ {
		if exit == true {
			break
		}
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
