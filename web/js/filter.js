document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('filter-creation-date').addEventListener('keyup', applyFilters);
    document.getElementById('filter-first-album-date').addEventListener('change', applyFilters);
    document.getElementById('filter-number-of-members').addEventListener('change', applyFilters);
    document.getElementById('filter-concert-locations').addEventListener('keyup', applyFilters);
    document.getElementById('reset-filters').addEventListener('mouseup', resetFilters);
});

function resetFilters(){
    document.getElementById('filter-creation-date').value = "";
    document.getElementById('filter-first-album-date').value = "";
    document.getElementById('filter-number-of-members').value = "";
    document.getElementById('filter-concert-locations').value = "";
    applyFilters();
}

function applyFilters() {
    fetch('/api/search/every/informations')
        .then(response => response.json())
        .then(data => {

            let creationDateInput = document.getElementById('filter-creation-date').value;
            let reverseFirstAlbumDateInput = document.getElementById('filter-first-album-date').value;
            let firstAlbumDateInput = reverseFirstAlbumDateInput.split("-").reverse().join("-");
            let numberOfMembersInput = parseInt(document.getElementById('filter-number-of-members').value);
            let concertLocationsInput = document.getElementById('filter-concert-locations').value.toLowerCase();

            let filteredResults = data.filter(artist => {

                let matchesCreationDate = false
                if (artist.creationDate.toString().includes(creationDateInput.toString())) {
                    matchesCreationDate = true
                }

                let matchesFirstAlbum = false
                if (artist.firstAlbum.includes(firstAlbumDateInput)) {
                    matchesFirstAlbum = true
                }

                let matchesNumberOfMembers = isNaN(numberOfMembersInput)
                if (artist.members.length === numberOfMembersInput) {
                    matchesNumberOfMembers = true
                }

                let matchesConcertLocation = false
                for (let i = 0; i < artist.locations.length; i++) {
                    if (artist.locations[i].toLowerCase().includes(concertLocationsInput)) {
                        matchesConcertLocation = true
                        break;
                    }
                }

                return matchesCreationDate && matchesFirstAlbum && matchesNumberOfMembers && matchesConcertLocation;
            });

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