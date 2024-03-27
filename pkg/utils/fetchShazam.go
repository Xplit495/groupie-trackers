package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"unicode"
)

/*type Singer struct { //Old structure made by us
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type TrackImage struct {
	Image string `json:"coverart"`
}

type Track struct {
	Title    string     `json:"title"`
	Subtitle string     `json:"subtitle"`
	Images   TrackImage `json:"images"`
}

type SingerHit struct {
	Singer Singer `json:"artist"`
}

type TrackHit struct {
	Track   Track  `json:"track"`
	Snippet string `json:"snippet"` //Old structure made by us
}

type SingersResponse struct {
	Hits []SingerHit `json:"hits"`
}

type TracksResponse struct {
	Hits []TrackHit `json:"hits"`
}

type ShazamResponse struct {
	Tracks  TracksResponse  `json:"tracks"`
	Singers SingersResponse `json:"artists"` //Old structure made by us
}*/

// ShazamResponse new structure simplify by ChatGPT

type ShazamResponse struct {
	Tracks struct {
		Hits []struct {
			Track struct {
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Images   struct {
					Image string `json:"coverart"`
				} `json:"images"`
				Hub struct {
					Options []struct {
						Actions []struct {
							Uri string `json:"uri"`
						} `json:"actions"`
					} `json:"options"`
				} `json:"hub"`
			} `json:"track"`
			Snippet string `json:"snippet"`
		} `json:"hits"`
	} `json:"tracks"`
	Singers struct {
		Hits []struct {
			Singer struct {
				Name   string `json:"name"`
				Avatar string `json:"avatar"`
			} `json:"artist"`
		} `json:"hits"`
	} `json:"artists"`
}

func FetchShazam(shazamInput string) []byte {

	cleanShazamInput := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return ' '
	}, shazamInput)

	url := fmt.Sprintf("https://shazam.p.rapidapi.com/search?term=%s&locale=fr-FR&offset=0&limit=6", cleanShazamInput)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	req.Header.Add("X-RapidAPI-Key", "f6c32c8c75msh764ee496b5cc231p1721f1jsn9e9dd8fb6363")
	req.Header.Add("X-RapidAPI-Host", "shazam.p.rapidapi.com")

	res, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		log.Fatal(err1)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
		}
	}(res.Body)

	data, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		log.Fatal(err3)
		return nil
	}

	fmt.Println()
	fmt.Println(cleanShazamInput)
	fmt.Println("=====================================")
	fmt.Println(res)
	fmt.Println("=====================================")
	fmt.Println(string(data))

	var shazamData ShazamResponse

	if err4 := json.Unmarshal(data, &shazamData); err4 != nil {
		log.Fatal(err4)
	}

	jsonShazam, err5 := json.Marshal(shazamData)
	if err5 != nil {
		log.Fatal(err5)
	}

	fmt.Println("=====================================")
	fmt.Println(shazamData)

	return jsonShazam
}
