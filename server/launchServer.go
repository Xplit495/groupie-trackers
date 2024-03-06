package server

import (
	"GroupieTrackers/pkg/utils"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type Artist struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func LaunchServer() {
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
		artists, err := utils.FetchArtists()
		if err != nil {
			http.Error(writer, "Failed to fetch artists", http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "gallery.html")))
		err1 := tmpl.Execute(writer, artists)
		if err1 != nil {
			return
		}
	})

	http.HandleFunc("/artists.html", func(writer http.ResponseWriter, request *http.Request) {
		artistIDStr := request.URL.Query().Get("id")
		artistID, err := strconv.Atoi(artistIDStr)
		if err != nil {
			http.Error(writer, "Invalid artist ID", http.StatusBadRequest)
			return
		}

		artist, err := utils.FetchArtistDetails(artistID)
		if err != nil {
			http.Error(writer, "Failed to fetch artist details", http.StatusInternalServerError)
			return
		}

		artist.MemberBis = strings.Join(artist.Members, ",\n")
		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "artists.html")))
		if err1 := tmpl.Execute(writer, artist); err1 != nil {
			http.Error(writer, "Failed to render artist details", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/api/search/artists", func(writer http.ResponseWriter, request *http.Request) {
		request.URL.Query().Get("query")

		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(Body io.ReadCloser) {
			err1 := Body.Close()
			if err1 != nil {

			}
		}(resp.Body)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if request.Method == http.MethodOptions {
			writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			writer.WriteHeader(http.StatusNoContent)
			return
		}

		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Content-Type", "application/json")

		_, err = writer.Write(body)
		if err != nil {
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
