package server

import (
	"GroupieTrackers/pkg/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

// fullArtists holds the list of all artists fetched from the API.
var fullArtists = utils.FetchArtists()

// Server starts the HTTP server.
func Server() {
	// Clear terminal screen
	utils.ClearTerminal()

	// Define web directory and create a file server handler
	webDir := filepath.Join("web")
	fileServer := http.FileServer(http.Dir(webDir))

	// Handle root endpoint
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			http.ServeFile(writer, request, filepath.Join(webDir, "html", "index.html"))
		} else {
			fileServer.ServeHTTP(writer, request)
		}
	})

	// Handle gallery endpoint
	http.HandleFunc("/gallery.html", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "gallery.html")))
		err := tmpl.Execute(writer, fullArtists)
		if err != nil {
			return
		}
	})

	// Handle artists endpoint
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

	// Handle shazamPage endpoint
	http.HandleFunc("/shazamPage.html", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, filepath.Join(webDir, "html", "shazamPage.html"))
	})

	// Handle shazamResults endpoint
	http.HandleFunc("/shazamResults", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, filepath.Join(webDir, "html", "shazamResults.html"))
	})

	// Handle searchPage endpoint
	http.HandleFunc("/searchPage", func(writer http.ResponseWriter, request *http.Request) {
		queryValue := request.URL.Query().Get("query")
		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "searchPage.html")))
		err2 := tmpl.Execute(writer, map[string]string{"QueryValue": queryValue})
		if err2 != nil {
			return
		}
	})

	// Handle shazam API endpoint
	http.HandleFunc("/api/shazam", func(writer http.ResponseWriter, request *http.Request) {
		shazamInput := request.URL.Query().Get("query")

		shazamData := utils.FetchShazam(shazamInput)

		writer.Header().Set("Content-Type", "application/json")
		_, err3 := writer.Write(shazamData)
		if err3 != nil {
			return
		}
	})

	// Handle deezer API endpoint
	http.HandleFunc("/api/deezer", func(writer http.ResponseWriter, request *http.Request) {
		artistIDStr := request.URL.Query().Get("artistId")
		var artistName string

		for _, artist := range fullArtists {
			if strconv.Itoa(artist.ID) == artistIDStr {
				artistName = artist.Name
				break
			}
		}

		deezerID := utils.FetchDeezer(artistName)

		writer.Header().Set("Content-Type", "application/json")
		_, err4 := writer.Write([]byte(deezerID))
		if err4 != nil {
			return
		}
	})

	// Handle search/every/information API endpoint
	http.HandleFunc("/api/search/every/informations", func(writer http.ResponseWriter, request *http.Request) {
		jsonData := utils.FetchLocations(fullArtists)

		writer.Header().Set("Content-Type", "application/json")
		_, err5 := writer.Write(jsonData)
		if err5 != nil {
			return
		}
	})

	// Handle search/locations API endpoint
	http.HandleFunc("/api/search/locations", func(writer http.ResponseWriter, request *http.Request) {
		groupID := request.URL.Query().Get("id")
		goodGroup := utils.FetchRelations(groupID)

		writer.Header().Set("Content-Type", "application/json")
		err4 := json.NewEncoder(writer).Encode(goodGroup.DatesLocations)
		if err4 != nil {
			fmt.Println("Error encoding JSON: ", err4)
			return
		}
	})

	// Open default web browser to the server URL
	err5 := utils.OpenBrowser("http://localhost:8080/")
	if err5 != nil {
		fmt.Println("Failed to open browser:", err5)
		fmt.Println("Please navigate to http://localhost:8080/ to view the web application.")
	}

	// Start the HTTP server
	err6 := http.ListenAndServe(":8080", nil)
	if err6 != nil {
		fmt.Println("Failed to start server:", err6)
		return
	}
}
