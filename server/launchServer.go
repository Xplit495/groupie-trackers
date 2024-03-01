package server

import (
	"GroupieTrackers/pkg/utils"
	"fmt"
	"net/http"
	"path/filepath"
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
