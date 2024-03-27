const params = new URLSearchParams(window.location.search);
const queryValue = params.get('query');

fetch(`/api/shazam?query=${queryValue}`)
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if (data.artists.hits.length > 0) {
            displayArtists(data.artists);
        }
        if (data.tracks.hits.length > 0) {
            displayTracks(data.tracks);
        }
    });

function displayArtists(artists) {
    const container = document.getElementById('artists-container');
    artists.hits.forEach(artistHit => {
        const artist = artistHit.artist;
        const element = document.createElement('div');
        element.innerHTML = `<h3>${artist.name}</h3><img src="${artist.avatar}" alt="${artist.name}" style="width: 100px; height: 100px;">`;
        container.appendChild(element);
    });
}

function displayTracks(tracks) {
    const container = document.getElementById('tracks-container');
    tracks.hits.forEach(trackHit => {
        const track = trackHit.track;
        const element = document.createElement('div');
        element.innerHTML = `<h3>${track.title} - ${track.subtitle}</h3><a href="${track.url}" target="_blank">Écouter</a><img src="${track.url}" alt="${track.name}" style="width: 100px; height: 100px;">`;
        container.appendChild(element);
    });
}