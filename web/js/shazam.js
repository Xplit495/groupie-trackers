document.addEventListener('DOMContentLoaded', function() {
    const params = new URLSearchParams(window.location.search);
    const queryValue = params.get('query');
    const loadingScreen = document.getElementById('loadingScreen');

    showLoadingScreen();

    fetch(`/api/shazam?query=${queryValue}`)
        .then(response => response.json())
        .then(data => {
            console.log(data);
            displayArtists(data.artists);
            displayTracks(data.tracks);
        });

    function displayArtists(artists) {
        hideLoadingScreen();
        const container = document.getElementById('artists-container');

        const textContainer = document.getElementById('text-artists-container')
        const text = document.createElement('h2');
        text.innerHTML = '<h2>Artiste(s) en lien trouvé :</h2>';
        text.style.color = 'white';
        textContainer.prepend(text);

        if (artists.hits === null) {
            const text = document.createElement('h2');
            text.innerHTML = '<h2>Aucun artiste trouvé</h2>';
            container.appendChild(text);
            return;
        }

        artists.hits.forEach(artistHit => {

            const artist = artistHit.artist;
            const artistContainer = document.createElement('div');
            artistContainer.id = 'artist-container';

            artistContainer.innerHTML = `<img src="${artist.avatar}" alt="${artist.name}"><h2>${artist.name}</h2>`;

            container.appendChild(artistContainer);
        });
    }

    function displayTracks(tracks) {

        const container = document.getElementById('tracks-container');

        const textContainer = document.getElementById('text-tracks-container')
        const text = document.createElement('h2');
        text.innerHTML = '<h2>Musique(s) en lien trouvé :</h2>';
        text.style.color = 'white';
        textContainer.prepend(text);

        if (tracks.hits === null) {
            const text = document.createElement('h2');
            text.innerHTML = '<h2>Aucune musique trouvée</h2>';
            container.appendChild(text);
            return;
        }

        tracks.hits.forEach(trackHit => {

            const track = trackHit.track;
            const trackContainer = document.createElement('div');
            trackContainer.id = 'track-container';

            trackContainer.innerHTML = `<h2>${track.title} - ${track.subtitle}</h2><a href="${track.hub.options[0].actions[0].uri}" target="_blank"><img src="${track.images.coverart}" alt="${track.title}"></a>`;

            container.appendChild(trackContainer);
        });
    }

    function showLoadingScreen() {
        loadingScreen.style.display = 'block';
    }

    function hideLoadingScreen() {
        loadingScreen.style.display = 'none';
    }

});
