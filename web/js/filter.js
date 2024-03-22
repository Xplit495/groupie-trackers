document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('filter-creation-date').addEventListener('keyup', applyFiltersAndDisplayResults);
    document.getElementById('filter-first-album-date').addEventListener('change', applyFiltersAndDisplayResults);
    document.getElementById('filter-number-of-members').addEventListener('change', applyFiltersAndDisplayResults);
    document.getElementById('filter-concert-locations').addEventListener('keyup', applyFiltersAndDisplayResults);

    ['filter-creation-date', 'filter-first-album-date', 'filter-number-of-members', 'filter-concert-locations'].forEach(id => {
        document.getElementById(id).addEventListener('keyup', function(event) {
            if (event.key === 'Enter') {
                applyFiltersAndDisplayResults();
            }
        });
    });
});

function applyFiltersAndDisplayResults() {
    fetch('/api/search/artists')
        .then(response => response.json())
        .then(data => {
            var creationDateFilter = document.getElementById('filter-creation-date').value;
            var reverseFirstAlbumDateFilter = document.getElementById('filter-first-album-date').value;
            var firstAlbumDateFilter = reverseFirstAlbumDateFilter.split("-").reverse().join("-");
            var numberOfMembersFilter = parseInt(document.getElementById('filter-number-of-members').value, 10);
            var concertLocationsFilter = document.getElementById('filter-concert-locations').value.toLowerCase();

            var filteredResults = data.filter(artist => {

                let matchesCreationDate = false
                if (artist.creationDate.toString().includes(creationDateFilter.toString())){
                    matchesCreationDate = true
                }

                let matchesFirstAlbum = false
                if (artist.firstAlbum.includes(firstAlbumDateFilter)){
                    matchesFirstAlbum = true
                }

                let matchesNumberOfMembers = isNaN(numberOfMembersFilter) || (artist.members && artist.members.length === numberOfMembersFilter);


                //let matchesConcertLocations = !concertLocationsFilter || (artist.concertLocations && artist.concertLocations.some(location => location.toLowerCase().includes(concertLocationsFilter)));

                return matchesCreationDate && matchesFirstAlbum && matchesNumberOfMembers //matchesConcertLocations;
            });

            console.log(filteredResults)
            displayFilteredResults(filteredResults);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function displayFilteredResults(filteredResults) {
    var cardsGrid = document.querySelector('.cards-grid');
    cardsGrid.innerHTML = '';

    filteredResults.forEach(artist => {
        var cardContainer = document.createElement('div');
        cardContainer.className = 'card-container';

        var cardLink = document.createElement('a');
        cardLink.href = `artists.html?id=${artist.id}`;

        var cardImageDiv = document.createElement('div');
        cardImageDiv.className = 'card-image';

        var image = document.createElement('img');
        image.src = artist.image;
        image.alt = `Photo de ${artist.name}`;

        var artistLine = document.createElement('div');
        artistLine.className = 'artist-line';

        var cardContent = document.createElement('div');
        cardContent.className = 'card-content';

        var artistName = document.createElement('p');
        artistName.className = 'artist-name';
        artistName.textContent = artist.name;

        cardImageDiv.appendChild(image);
        cardContent.appendChild(artistName);

        cardLink.appendChild(cardImageDiv);
        cardLink.appendChild(artistLine);
        cardLink.appendChild(cardContent);

        cardContainer.appendChild(cardLink);
        cardsGrid.appendChild(cardContainer);
    });
}