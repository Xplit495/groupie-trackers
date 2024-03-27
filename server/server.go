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

var fullArtists = utils.FetchArtists()

func Server() {
	utils.ClearTerminal()

	webDir := filepath.Join("web")
	fileServer := http.FileServer(http.Dir(webDir))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			http.ServeFile(writer, request, filepath.Join(webDir, "html", "index.html"))
		} else {
			fileServer.ServeHTTP(writer, request)
		}
	})

	http.HandleFunc("/gallery.html", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "gallery.html")))

		err1 := tmpl.Execute(writer, fullArtists)
		if err1 != nil {
			return
		}
	})

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

	http.HandleFunc("/shazamPage.html", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, filepath.Join(webDir, "html", "shazamPage.html"))

	})

	http.HandleFunc("/shazamResults", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, filepath.Join(webDir, "html", "shazamResults.html"))
	})

	http.HandleFunc("/searchPage", func(writer http.ResponseWriter, request *http.Request) {
		queryValue := request.URL.Query().Get("query")

		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "searchPage.html")))

		err := tmpl.Execute(writer, map[string]string{"QueryValue": queryValue})
		if err != nil {
			return
		}
	})

	http.HandleFunc("/api/shazam", func(writer http.ResponseWriter, request *http.Request) {
		shazamInput := request.URL.Query().Get("query")

		shazamData := utils.FetchShazam(shazamInput)

		_, err1 := writer.Write(shazamData)
		if err1 != nil {
			return
		}

	})

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

		_, err1 := writer.Write([]byte(deezerID))
		if err1 != nil {
			return
		}

	})

	http.HandleFunc("/api/search/every/informations", func(writer http.ResponseWriter, request *http.Request) {
		jsonData := utils.FetchLocations(fullArtists)

		writer.Header().Set("Content-Type", "application/json")

		_, err1 := writer.Write(jsonData)
		if err1 != nil {
			return
		}

	})

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
