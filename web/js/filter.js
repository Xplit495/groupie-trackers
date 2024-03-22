document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('filter-creation-date').addEventListener('keyup', applyFiltersAndDisplayResults);
    document.getElementById('filter-first-album-date').addEventListener('change', applyFiltersAndDisplayResults);
    document.getElementById('filter-number-of-members').addEventListener('change', applyFiltersAndDisplayResults);
    document.getElementById('filter-concert-locations').addEventListener('keyup', applyFiltersAndDisplayResults);
    document.getElementById('reset-filters').addEventListener('mouseup', applyFiltersAndDisplayResults);

    ['filter-creation-date', 'filter-first-album-date', 'filter-number-of-members', 'filter-concert-locations'].forEach(id => {
        document.getElementById(id).addEventListener('keyup', function(event) {
            if (event.key === 'Enter') {
                applyFiltersAndDisplayResults();
            }
        });
    });
});

document.addEventListener('DOMContentLoaded', function () {
    let resetButton = document.getElementById('reset-filters');

    resetButton.addEventListener('click', function () {
        document.getElementById('filter-creation-date').value = "";
        document.getElementById('filter-first-album-date').value = "";
        document.getElementById('filter-number-of-members').value = "";
        document.getElementById('filter-concert-locations').value = "";
    });
});

function applyFiltersAndDisplayResults() {
    fetch('/api/search/artists')
        .then(response => response.json())
        .then(data => {
            let creationDateFilter = document.getElementById('filter-creation-date').value;
            let reverseFirstAlbumDateFilter = document.getElementById('filter-first-album-date').value;
            let firstAlbumDateFilter = reverseFirstAlbumDateFilter.split("-").reverse().join("-");
            let numberOfMembersFilter = parseInt(document.getElementById('filter-number-of-members').value, 10);
            let concertLocationsFilter = document.getElementById('filter-concert-locations').value.toLowerCase();

            let filteredResults = data.filter(artist => {

                let matchesCreationDate = false
                if (artist.creationDate.toString().includes(creationDateFilter.toString())) {
                    matchesCreationDate = true
                }

                let matchesFirstAlbum = false
                if (artist.firstAlbum.includes(firstAlbumDateFilter)) {
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
    let cardsGrid = document.querySelector('.cards-grid');
    cardsGrid.innerHTML = '';

    filteredResults.forEach(artist => {
        let cardContainer = document.createElement('div');
        cardContainer.className = 'card-container';

        let cardLink = document.createElement('a');
        cardLink.href = `artists.html?id=${artist.id}`;

        let cardImageDiv = document.createElement('div');
        cardImageDiv.className = 'card-image';

        let image = document.createElement('img');
        image.src = artist.image;
        image.alt = `Photo de ${artist.name}`;

        let artistLine = document.createElement('div');
        artistLine.className = 'artist-line';

        let cardContent = document.createElement('div');
        cardContent.className = 'card-content';

        let artistName = document.createElement('p');
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