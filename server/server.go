package server

import (
	"GroupieTrackers/pkg/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
)

type Location struct {
	ID        int      `json:"id"`
	Image     string   `json:"image"`
	Locations []string `json:"locations"`
}

type Locations struct {
	Index []Location `json:"index"`
}

var fullArtists = utils.FetchArtists()

func Server() {
	utils.ClearTerminal()

	//PAGE D'ACCUEIL

	webDir := filepath.Join("web")
	fileServer := http.FileServer(http.Dir(webDir))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			http.ServeFile(writer, request, filepath.Join(webDir, "html", "index.html"))
		} else {
			fileServer.ServeHTTP(writer, request)
		}
	})

	//Galerie d'artistes

	http.HandleFunc("/gallery.html", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "gallery.html")))

		err1 := tmpl.Execute(writer, fullArtists)
		if err1 != nil {
			return
		}

	})

	//Page d'artiste

	http.HandleFunc("/artists.html", func(writer http.ResponseWriter, request *http.Request) {
		artistIDStr := request.URL.Query().Get("id")

		var artist utils.Artist

		for i := 0; i < len(fullArtists); i++ {
			if strconv.Itoa(fullArtists[i].ID) == artistIDStr {
				artist = fullArtists[i]
			}
		}

		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "artists.html")))

		if err1 := tmpl.Execute(writer, artist); err1 != nil {
			http.Error(writer, "Failed to render artist details", http.StatusInternalServerError)
		}

	})

	//Barre de recherche artistes

	http.HandleFunc("/api/search/artists", func(writer http.ResponseWriter, request *http.Request) {

		fullArtistsJson, err := json.Marshal(fullArtists)
		if err != nil {
			fmt.Println("Error serializing artists to JSON: ", err)
		}

		writer.Header().Set("Content-Type", "application/json")

		_, err1 := writer.Write(fullArtistsJson)
		if err1 != nil {
			return
		}
	})

	//Barre de recherche lieux de concerts

	http.HandleFunc("/api/search/locations/search/bar", func(writer http.ResponseWriter, request *http.Request) {
		resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Failed to close response body:", err)
			}
		}(resp.Body)

		data, _ := io.ReadAll(resp.Body)

		var response Locations

		err := json.Unmarshal(data, &response)
		if err != nil {
			fmt.Println("Error parsing JSON: ", err)
			return
		}

		for i, location := range response.Index {
			for _, artist := range fullArtists {
				if location.ID == artist.ID {
					response.Index[i].Image = artist.Image
					break
				}
			}
		}

		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error serializing map to JSON: ", err)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		_, err1 := writer.Write(jsonData)
		if err1 != nil {
			return
		}

	})

	//MAPS

	http.HandleFunc("/api/search/locations", func(writer http.ResponseWriter, request *http.Request) {
		groupID := request.URL.Query().Get("id")

		goodGroup := utils.FetchRelations(groupID)

		writer.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(writer).Encode(goodGroup.DatesLocations)
		if err != nil {
			fmt.Println("Error encoding JSON: ", err)
			return
		}
	})

	//LANCE LE SERVEUR

	err := utils.OpenBrowser("http://localhost:8080/")
	if err != nil {
		fmt.Println("Failed to open browser:", err)
		return
	}

	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		fmt.Println("Failed to start server:", err1)
		return
	}
}
