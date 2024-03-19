document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('search-input').addEventListener('keyup', searchArtist);
    document.getElementById('filter-creation-date').addEventListener('change', searchArtist);
    document.getElementById('filter-first-album-date').addEventListener('change', searchArtist);
    document.getElementById('filter-number-of-members').addEventListener('change', searchArtist);
    document.getElementById('filter-concert-locations').addEventListener('change', searchArtist);
    document.addEventListener('click', function(event) {
        var isClickInsideSearchBox = document.getElementById('search-box').contains(event.target);
        var isClickInsideSuggestions = document.getElementById('suggestions').contains(event.target);

        if (!isClickInsideSearchBox && !isClickInsideSuggestions) {
            document.getElementById('suggestions').style.display = 'none';
        }
    });
});

function searchArtist() {
    var input = document.getElementById('search-input').value.toLowerCase();
    var creationDateFilter = document.getElementById('filter-creation-date').value;
    var firstAlbumDateFilter = document.getElementById('filter-first-album-date').value;
    var numberOfMembersFilter = parseInt(document.getElementById('filter-number-of-members').value, 10);
    var concertLocationsFilter = document.getElementById('filter-concert-locations').value.toLowerCase();

    if (input.length > 0 || creationDateFilter || firstAlbumDateFilter || numberOfMembersFilter || concertLocationsFilter) {
        fetch('/api/search/artists?query=' + encodeURIComponent(input))
            .then(response => response.json())
            .then(data => {
                var suggestions = data.filter(artist => {
                    let matchesName = artist.name.toLowerCase().includes(input);
                    let matchesCreationDate = creationDateFilter ? artist.creationDate === creationDateFilter : true;
                    let matchesFirstAlbumDate = firstAlbumDateFilter ? artist.firstAlbumDate === firstAlbumDateFilter : true;
                    let matchesNumberOfMembers = numberOfMembersFilter ? artist.members.length === numberOfMembersFilter : true;
                    let matchesConcertLocations = concertLocationsFilter ? artist.concertLocations.some(location => location.toLowerCase().includes(concertLocationsFilter)) : true;

                    return matchesName && matchesCreationDate && matchesFirstAlbumDate && matchesNumberOfMembers && matchesConcertLocations;
                });

                showSuggestions(suggestions);
            })
            .catch(error => {
                console.error('Error:', error);
                document.getElementById('suggestions').style.display = 'none';
            });
    } else {
        document.getElementById('suggestions').style.display = 'none';
    }
}

function showSuggestions(suggestions) {
    var suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';

    if (suggestions.length > 0) {
        suggestions.forEach(artist => {
            var suggestionElement = document.createElement('div');
            suggestionElement.classList.add('suggestion-item');

            var artistImage = document.createElement('img');
            artistImage.src = artist.image;
            artistImage.alt = artist.name;
            artistImage.classList.add('artist-image');

            var artistName = document.createElement('p');
            artistName.textContent = artist.name;
            artistName.classList.add('name-artist');

            suggestionElement.appendChild(artistImage);
            suggestionElement.appendChild(artistName);

            suggestionElement.addEventListener('click', function() {
                window.location.href = "/artists.html?id=" + artist.id;
            });

            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}