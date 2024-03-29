package server

import (
	"GroupieTrackers/pkg/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// fullArtists holds the list of all artists fetched from the API.
var fullArtists = utils.FetchArtists()
var formatted = "02/01/2006 15:04:05"

// Server starts the HTTP server.
func Server() {
	// Open log file
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture du fichier de log:", err)
	}
	defer func(logFile *os.File) {
		err1 := logFile.Close()
		if err1 != nil {

		}
	}(logFile)

	log.SetOutput(logFile)
	log.SetFlags(0)

	// Clear terminal screen
	utils.ClearTerminal()

	// Define web directory and create a file server handler
	webDir := filepath.Join("web")
	fileServer := http.FileServer(http.Dir(webDir))

	// Handle root endpoint
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {

			ip := request.Header.Get("X-Forwarded-For")
			if ip == "" {
				ip = request.RemoteAddr
			}

			fmt.Print("Accueil // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")
			log.Print("Accueil // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")

			//No cache for logs
			writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
			writer.Header().Set("Pragma", "no-cache")
			writer.Header().Set("Expires", "0")
			//No cache for logs

			http.ServeFile(writer, request, filepath.Join(webDir, "html", "index.html"))
		} else {
			fileServer.ServeHTTP(writer, request)
		}
	})

	// Handle gallery endpoint
	http.HandleFunc("/gallery.html", func(writer http.ResponseWriter, request *http.Request) {
		ip := request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = request.RemoteAddr
		}

		fmt.Print("Galerie // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")
		log.Print("Galerie // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")

		//No cache for logs
		writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		writer.Header().Set("Pragma", "no-cache")
		writer.Header().Set("Expires", "0")
		//No cache for logs

		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "gallery.html")))
		err2 := tmpl.Execute(writer, fullArtists)
		if err2 != nil {
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

		ip := request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = request.RemoteAddr
		}

		fmt.Print("Artist : '", artist.Name, "' // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")
		log.Print("Artist : '", artist.Name, "' // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")

		//No cache for logs
		writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		writer.Header().Set("Pragma", "no-cache")
		writer.Header().Set("Expires", "0")
		//No cache for logs

		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "artists.html")))
		if err3 := tmpl.Execute(writer, artist); err3 != nil {
			http.Error(writer, "Failed to render artist details", http.StatusInternalServerError)
		}
	})

	// Handle shazamPage endpoint
	http.HandleFunc("/shazamPage.html", func(writer http.ResponseWriter, request *http.Request) {
		ip := request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = request.RemoteAddr
		}

		fmt.Print("Shazam // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")
		log.Print("Shazam // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")

		//No cache for logs
		writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		writer.Header().Set("Pragma", "no-cache")
		writer.Header().Set("Expires", "0")
		//No cache for logs

		http.ServeFile(writer, request, filepath.Join(webDir, "html", "shazamPage.html"))
	})

	// Handle shazamResults endpoint
	http.HandleFunc("/shazamResults", func(writer http.ResponseWriter, request *http.Request) {
		ip := request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = request.RemoteAddr
		}

		fmt.Print("Page Résultats Shazam // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n")
		log.Print("Page Résultats Shazam // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n")

		//No cache for logs
		writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		writer.Header().Set("Pragma", "no-cache")
		writer.Header().Set("Expires", "0")
		//No cache for logs

		http.ServeFile(writer, request, filepath.Join(webDir, "html", "shazamResults.html"))
	})

	// Handle searchPage endpoint
	http.HandleFunc("/searchPage", func(writer http.ResponseWriter, request *http.Request) {
		queryValue := request.URL.Query().Get("query")
		ip := request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = request.RemoteAddr
		}

		fmt.Print("Page Résultats Recherche // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n")
		fmt.Print("Recherche : '", queryValue, "' // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")
		log.Print("Page Résultats Recherche // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n")
		log.Print("Recherche : '", queryValue, "' // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")

		//No cache for logs
		writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		writer.Header().Set("Pragma", "no-cache")
		writer.Header().Set("Expires", "0")
		//No cache for logs

		tmpl := template.Must(template.ParseFiles(filepath.Join(webDir, "html", "searchPage.html")))
		err4 := tmpl.Execute(writer, map[string]string{"QueryValue": queryValue})
		if err4 != nil {
			return
		}
	})

	// Handle shazam API endpoint
	http.HandleFunc("/api/shazam", func(writer http.ResponseWriter, request *http.Request) {
		shazamInput := request.URL.Query().Get("query")

		ip := request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = request.RemoteAddr
		}

		fmt.Print("Recherche Shazam : ", shazamInput, " // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")
		log.Print("Recherche Shazam : ", shazamInput, " // (Date : ", time.Now().Format(formatted), ") (IP : ", ip, ")\n\n")

		shazamData := utils.FetchShazam(shazamInput)

		writer.Header().Set("Content-Type", "application/json")
		_, err5 := writer.Write(shazamData)
		if err5 != nil {
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
		_, err6 := writer.Write([]byte(deezerID))
		if err6 != nil {
			return
		}
	})

	// Handle search/every/information API endpoint
	http.HandleFunc("/api/search/every/informations", func(writer http.ResponseWriter, request *http.Request) {
		jsonData := utils.FetchLocations(fullArtists)

		writer.Header().Set("Content-Type", "application/json")
		_, err7 := writer.Write(jsonData)
		if err7 != nil {
			return
		}
	})

	// Handle search/locations API endpoint
	http.HandleFunc("/api/search/locations", func(writer http.ResponseWriter, request *http.Request) {
		groupID := request.URL.Query().Get("id")
		goodGroup := utils.FetchRelations(groupID)

		writer.Header().Set("Content-Type", "application/json")
		err6 := json.NewEncoder(writer).Encode(goodGroup.DatesLocations)
		if err6 != nil {
			fmt.Println("Error encoding JSON: ", err6)
			return
		}
	})

	// Open default web browser to the server URL
	err7 := utils.OpenBrowser("http://localhost:8080/")
	if err7 != nil {
		fmt.Println("Failed to open browser:", err7)
		fmt.Println("Please navigate to http://localhost:8080/ to view the web application.")
	}

	// Start the HTTP server
	err8 := http.ListenAndServe(":8080", nil)
	if err8 != nil {
		fmt.Println("Failed to start server:", err8)
		return
	}
}
