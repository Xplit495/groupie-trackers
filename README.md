
# Groupie-Trackers

Groupie-Trackers is a project aimed at training individuals to use APIs.


## Appendix

**!!! The header search bar and Shazam search bar are completely different !!!**


## Usage/Examples

Open a terminal into the project file and write :

```
go run .\cmd\main.go
```

After that, your browser needs to automatically open it on the correct page.

Else go to : ```http://localhost:8080/```

# Features

- Dynamic search bar (in the header):
    - Functions with: concert venues, artist names, creation dates, first albums, and members

- ## Gallery
    - References every artist provided by the main "Groupie-Trackers" API
    - All filters are available and functional
    - Reset button also functions properly
    - Each artist card is clickable

- ## Artists
    - Each artist's page provides detailed information about them/us: creation date, name, first album, and member(s).
    - Deezer API is also included to play songs and discover artists/groups
    - Google Maps API is integrated to view concert locations directly on the website, with watermarks pinned on the map indicating concert dates
- ## Shazam
    - Shazam is an additional feature that works similarly to Shazam but with more capabilities.
    - Every music cover is clickable and redirect to AppleMusic to stream.
    - For example: You can input a name, lyrics, or anything else related to any artist (even if they are not in the original Groupie Trackers API), and receive multiple pieces of information related to your search input.


    
## Tech Stack

**API**: Deezer, GoogleMaps, AppleMusic, Shazam (By Api Dojo on https://rapidapi.com/apidojo/api/shazam)

**Client:** HTML/CSS, JavaScript

**Server:** Golang


## Authors

- Carrola Quentin
- Malagouen Shemsedine

## Gitea

- https://ytrack.learn.ynov.com/git/carquentin/GroupieTrackers

