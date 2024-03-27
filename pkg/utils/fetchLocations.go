package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Location represents information about a location.
type Location struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string `json:"locations"`
}

// FetchLocations retrieves location data from the remote API and updates it with additional artist information.
func FetchLocations(fullArtists []Artist) []byte {
	// Fetch location data from the remote API
	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	// Ensure response body is closed when function returns
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Failed to close response body:", err)
		}
	}(resp.Body)

	// Read the response body
	data, _ := io.ReadAll(resp.Body)

	// Temporary struct to hold decoded JSON data
	var temp struct {
		Index []Location `json:"index"`
	}

	// Unmarshal JSON data into the temporary struct
	if err1 := json.Unmarshal(data, &temp); err1 != nil {
		fmt.Println("Error decoding:", err1)
		return nil
	}

	// Final data containing location information with updated artist details
	finalData := temp.Index

	// Update location information with artist details
	for i, location := range finalData {
		for _, artist := range fullArtists {
			if location.ID == artist.ID {
				finalData[i].Name = artist.Name
				finalData[i].Image = artist.Image
				finalData[i].Members = artist.Members
				finalData[i].CreationDate = artist.CreationDate
				finalData[i].FirstAlbum = artist.FirstAlbum
				break
			}
		}
	}

	// Serialize final data to JSON
	jsonData, err2 := json.Marshal(finalData)
	if err2 != nil {
		fmt.Println("Error serializing map to JSON: ", err2)
		return nil
	}

	return jsonData
}
