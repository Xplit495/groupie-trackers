document.addEventListener('DOMContentLoaded', function () {
    // Attacher les événements de changement pour appliquer les filtres directement
    document.getElementById('filter-creation-date').addEventListener('change', applyFiltersAndDisplayResults);
    document.getElementById('filter-first-album-date').addEventListener('change', applyFiltersAndDisplayResults);
    document.getElementById('filter-number-of-members').addEventListener('change', applyFiltersAndDisplayResults);
    document.getElementById('filter-concert-locations').addEventListener('change', applyFiltersAndDisplayResults);
});
function applyFiltersAndDisplayResults() {
    // Exemple de récupération de données filtrées, adaptez selon votre backend/API
    fetch('/api/search/artists')
        .then(response => response.json())
        .then(data => {
            var creationDateFilter = document.getElementById('filter-creation-date').value;
            var firstAlbumDateFilter = document.getElementById('filter-first-album-date').value;
            var numberOfMembersFilter = parseInt(document.getElementById('filter-number-of-members').value, 10);
            var concertLocationsFilter = document.getElementById('filter-concert-locations').value.toLowerCase();
            var filteredResults = data.filter(artist => {
                let matchesCreationDate = !creationDateFilter || new Date(artist.creationDate) >= new Date(creationDateFilter);
                let matchesFirstAlbumDate = !firstAlbumDateFilter || new Date(artist.firstAlbumDate) >= new Date(firstAlbumDateFilter);
                let matchesNumberOfMembers = isNaN(numberOfMembersFilter) || (artist.members && artist.members.length === numberOfMembersFilter);
                let matchesConcertLocations = !concertLocationsFilter || (artist.concertLocations && artist.concertLocations.some(location => location.toLowerCase().includes(concertLocationsFilter)));
                return matchesCreationDate && matchesFirstAlbumDate && matchesNumberOfMembers && matchesConcertLocations;
            });
            displayFilteredResults(filteredResults);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}
function displayFilteredResults(filteredResults) {
    var cardsGrid = document.querySelector('.cards-grid');
    cardsGrid.innerHTML = ''; // Efface les résultats précédents
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