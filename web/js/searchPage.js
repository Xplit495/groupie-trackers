// Extracting query parameter from the URL
const params = new URLSearchParams(window.location.search);
const queryValue = params.get('query');

// Checking if query parameter exists and is not empty
if (queryValue && queryValue.length > 0) {
    let suggestions = [];

    // Fetching data from API for search
    fetch('/api/search/every/informations')
        .then(response => response.json())
        .then(data => {
            // Iterating through fetched data to find matching suggestions
            data.forEach(band => {
                if (band.name.toLowerCase().includes(queryValue.toLowerCase())) {
                    suggestions.push({
                        name: band.name,
                        type: 'Groupe',
                        image: band.image,
                        id: band.id
                    });
                }

                band.members.forEach(member => {
                    if (member.toLowerCase().includes(queryValue.toLowerCase())) {
                        suggestions.push({
                            name: member,
                            type: 'Membre',
                            image: band.image,
                            id: band.id
                        });
                    }
                });

                if (band.firstAlbum && band.firstAlbum.toLowerCase().includes(queryValue.toLowerCase())) {
                    suggestions.push({
                        name: band.firstAlbum,
                        type: 'FirstAlbum',
                        image: band.image,
                        id: band.id
                    });
                }

                if (band.creationDate && band.creationDate.toString().toLowerCase().includes(queryValue.toLowerCase())) {
                    suggestions.push({
                        name: band.creationDate.toString(),
                        type: 'CreationDate',
                        image: band.image,
                        id: band.id
                    });
                }

                band.locations.forEach(location => {
                    if (location.toLowerCase().includes(queryValue.toLowerCase())) {
                        suggestions.push({
                            name: location,
                            type: 'Location',
                            image: band.image,
                            id: band.id,
                        });
                    }
                });

            });

            // Handling case where no results are found
            if (suggestions.length === 0) {
                let noResults = document.createElement('p');
                noResults.textContent = 'Aucun résultat trouvé';
                noResults.style.textAlign = 'center';
                noResults.style.marginTop = '3%';
                noResults.style.fontSize = '200%';
                noResults.style.fontWeight = 'bold';
                document.querySelector('.cards-grid').appendChild(noResults);
            } else {
                // Displaying the search results
                displayedResults(suggestions);
            }
        })
        .catch(error => console.error('Error:', error));
}

// Function to display search results
function displayedResults(filteredResults) {
    let cardsGrid = document.querySelector('.cards-grid');
    cardsGrid.innerHTML = '';

    filteredResults.forEach(artist => {
        // Creating elements for displaying search results
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

        // Constructing artist name based on type
        if (artist.type === "Membre") {
            artistName.textContent = artist.name + " (Membre)";
        } else if (artist.type === "Groupe") {
            artistName.textContent = artist.name + " (Groupe)";
        } else if (artist.type === "FirstAlbum") {
            artistName.textContent = artist.name + " (Premier Album)";
        } else if (artist.type === "CreationDate") {
            artistName.textContent = artist.name + " (Date de Création)";
        } else if (artist.type === "Location") {
            artistName.textContent = artist.name + " (Lieu)";
        }

        // Appending elements to construct search result card
        cardImageDiv.appendChild(image);
        cardContent.appendChild(artistName);
        cardLink.appendChild(cardImageDiv);
        cardLink.appendChild(artistLine);
        cardLink.appendChild(cardContent);
        cardContainer.appendChild(cardLink);
        cardsGrid.appendChild(cardContainer);
    });
}
