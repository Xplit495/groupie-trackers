// Event listener to execute code when the DOM content is loaded
document.addEventListener('DOMContentLoaded', function () {
    // Adding event listeners to filter input elements
    document.getElementById('filter-creation-date').addEventListener('keyup', applyFilters);
    document.getElementById('filter-first-album-date').addEventListener('change', applyFilters);
    document.getElementById('filter-number-of-members').addEventListener('change', applyFilters);
    document.getElementById('filter-concert-locations').addEventListener('keyup', applyFilters);
    // Adding event listener to reset filters button
    document.getElementById('reset-filters').addEventListener('mouseup', resetFilters);
});

// Function to reset all filters to their default values
function resetFilters() {
    // Resetting filter input values to empty strings
    document.getElementById('filter-creation-date').value = "";
    document.getElementById('filter-first-album-date').value = "";
    document.getElementById('filter-number-of-members').value = "";
    document.getElementById('filter-concert-locations').value = "";
    // Applying filters with reset values
    applyFilters();
}

// Function to apply filters based on user input
function applyFilters() {
    // Fetching data from an API endpoint
    fetch('/api/search/every/informations')
        .then(response => response.json()) // Parsing the response as JSON
        .then(data => {
            // Retrieving filter input values
            let creationDateInput = document.getElementById('filter-creation-date').value;
            let reverseFirstAlbumDateInput = document.getElementById('filter-first-album-date').value;
            let firstAlbumDateInput = reverseFirstAlbumDateInput.split("-").reverse().join("-");
            let numberOfMembersInput = parseInt(document.getElementById('filter-number-of-members').value);
            let concertLocationsInput = document.getElementById('filter-concert-locations').value.toLowerCase();

            // Filtering the data based on user input
            let filteredResults = data.filter(artist => {
                // Checking if artist matches the filter conditions
                let matchesCreationDate = false;
                if (artist.creationDate.toString().includes(creationDateInput.toString())) {
                    matchesCreationDate = true;
                }

                let matchesFirstAlbum = false;
                if (artist.firstAlbum.includes(firstAlbumDateInput)) {
                    matchesFirstAlbum = true;
                }

                let matchesNumberOfMembers = isNaN(numberOfMembersInput);
                if (artist.members.length === numberOfMembersInput) {
                    matchesNumberOfMembers = true;
                }

                let matchesConcertLocation = false;
                for (let i = 0; i < artist.locations.length; i++) {
                    if (artist.locations[i].toLowerCase().includes(concertLocationsInput)) {
                        matchesConcertLocation = true;
                        break;
                    }
                }

                return matchesCreationDate && matchesFirstAlbum && matchesNumberOfMembers && matchesConcertLocation;
            });

            // Displaying filtered results
            displayFilteredResults(filteredResults);
        })
        .catch(error => {
            // Error handling in case the fetch operation fails
            console.error('Error:', error);
        });
}

// Function to display filtered results on the webpage
function displayFilteredResults(filteredResults) {
    // Finding the container for displaying artist cards
    let cardsGrid = document.querySelector('.cards-grid');
    // Clearing previous content
    cardsGrid.innerHTML = '';

    // Iterating through filtered results
    filteredResults.forEach(artist => {
        // Creating elements for displaying artist information
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

        // Appending elements to construct artist card
        cardImageDiv.appendChild(image);
        cardContent.appendChild(artistName);
        cardLink.appendChild(cardImageDiv);
        cardLink.appendChild(artistLine);
        cardLink.appendChild(cardContent);
        cardContainer.appendChild(cardLink);
        cardsGrid.appendChild(cardContainer);
    });
}
