package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	ID        int      `json:"id"`
	Image     string   `json:"image"`
	Locations []string `json:"locations"`
}

func FetchLocations(fullArtists []Artist) []byte {

	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body:", err)
		}
	}(resp.Body)

	data, _ := io.ReadAll(resp.Body)

	var temp struct {
		Index []Location `json:"index"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		fmt.Println("Erreur lors du décodage:", err)
		return nil
	}

	finalData := temp.Index

	for i, location := range finalData {
		for _, artist := range fullArtists {
			if location.ID == artist.ID {
				finalData[i].Image = artist.Image
				break
			}
		}
	}

	jsonData, err := json.Marshal(finalData)
	if err != nil {
		fmt.Println("Error serializing map to JSON: ", err)
		return nil
	}

	return jsonData
}
