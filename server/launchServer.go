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
	"strings"
)

type Artist struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Relation `json:"index"`
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

		resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Failed to close response body:", err)
			}
		}(resp.Body)

		data, _ := io.ReadAll(resp.Body)

		writer.Header().Set("Content-Type", "application/json")

		_, err := writer.Write(data)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/api/search/locations", func(writer http.ResponseWriter, request *http.Request) {
		groupID := request.URL.Query().Get("id")
		if groupID == "" {
			http.Error(writer, "ID de groupe manquant", http.StatusBadRequest)
			return
		}

		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
		if err != nil {
			http.Error(writer, fmt.Sprintf("Erreur lors de la requête à l'API: %v", err), http.StatusInternalServerError)
			return
		}
		defer func(Body io.ReadCloser) {
			err1 := Body.Close()
			if err1 != nil {

			}
		}(resp.Body)

		body, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			http.Error(writer, fmt.Sprintf("Erreur lors de la lecture de la réponse: %v", err), http.StatusInternalServerError)
			return
		}

		var relations Relations
		if err2 := json.Unmarshal(body, &relations); err2 != nil {
			http.Error(writer, fmt.Sprintf("Erreur lors du décodage du JSON: %v", err), http.StatusInternalServerError)
			return
		}

		intID, err3 := strconv.Atoi(groupID)
		if err3 != nil {
			http.Error(writer, fmt.Sprintf("Erreur lors de la conversion de l'ID en entier: %v", err), http.StatusBadRequest)
			return
		}

		for _, relation := range relations.Index {
			if relation.ID == intID {
				writer.Header().Set("Content-Type", "application/json")
				err2 := json.NewEncoder(writer).Encode(relation.DatesLocations)
				if err2 != nil {
					return
				}
				return
			}
		}

		http.Error(writer, "Groupe non trouvé", http.StatusNotFound)
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
