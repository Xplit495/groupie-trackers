document.addEventListener('DOMContentLoaded', function() {
    // Extracting query parameter from URL
    const params = new URLSearchParams(window.location.search);
    const queryValue = params.get('query');
    const loadingScreen = document.getElementById('loadingScreen');

    // Display loading screen
    showLoadingScreen();

    // Fetching data from API based on query
    fetch(`/api/shazam?query=${queryValue}`)
        .then(response => response.json())
        .then(data => {
            // Displaying artists and tracks after fetching data
            displayArtists(data.artists);
            displayTracks(data.tracks);
        });

    // Function to display artists
    function displayArtists(artists) {
        hideLoadingScreen();
        const container = document.getElementById('artists-container');

        const textContainer = document.getElementById('text-artists-container');
        const text = document.createElement('h2');
        text.innerHTML = '<h2>Artiste(s) en lien trouvé :</h2>'; // Header for artists
        text.style.color = 'white';
        textContainer.prepend(text);

        // If no artist found
        if (artists.hits === null) {
            const text = document.createElement('h2');
            text.innerHTML = '<h2>Aucun artiste trouvé</h2>'; // Display message for no artist found
            container.appendChild(text);
            return;
        }

        // Looping through each artist found
        artists.hits.forEach(artistHit => {
            const artist = artistHit.artist;
            const artistContainer = document.createElement('div');
            artistContainer.id = 'artist-container';

            // Displaying artist's image and name
            artistContainer.innerHTML = `<img src="${artist.avatar}" alt="${artist.name}"><h2>${artist.name}</h2>`;

            container.appendChild(artistContainer);
        });
    }

    // Function to display tracks
    function displayTracks(tracks) {
        const container = document.getElementById('tracks-container');

        const textContainer = document.getElementById('text-tracks-container');
        const text = document.createElement('h2');
        text.innerHTML = '<h2>Musique(s) en lien trouvé :</h2>'; // Header for tracks
        text.style.color = 'white';
        textContainer.prepend(text);

        // If no track found
        if (tracks.hits === null) {
            const text = document.createElement('h2');
            text.innerHTML = '<h2>Aucune musique trouvée</h2>'; // Display message for no track found
            container.appendChild(text);
            return;
        }

        // Looping through each track found
        tracks.hits.forEach(trackHit => {
            const track = trackHit.track;
            const trackContainer = document.createElement('div');
            trackContainer.id = 'track-container';

            // Displaying track's title, subtitle, and cover image
            trackContainer.innerHTML = `<h2>${track.title} - ${track.subtitle}</h2><a href="${track.hub.options[0].actions[0].uri}" target="_blank"><img src="${track.images.coverart}" alt="${track.title}"></a>`;

            container.appendChild(trackContainer);
        });
    }

    // Function to show loading screen
    function showLoadingScreen() {
        loadingScreen.style.display = 'block';
    }

    // Function to hide loading screen
    function hideLoadingScreen() {
        loadingScreen.style.display = 'none';
    }
});
