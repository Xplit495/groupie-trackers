package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string `json:"locations"`
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
				finalData[i].Name = artist.Name
				finalData[i].Image = artist.Image
				finalData[i].Members = artist.Members
				finalData[i].CreationDate = artist.CreationDate
				finalData[i].FirstAlbum = artist.FirstAlbum
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
