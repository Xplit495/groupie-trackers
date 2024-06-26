// Event listener to close the suggestions dropdown when a click occurs outside the search box or suggestions
document.addEventListener('click', function(event) {
    let isClickInsideSearchBox = document.getElementById('search-box').contains(event.target);
    let isClickInsideSuggestions = document.getElementById('suggestions').contains(event.target);

    if (!isClickInsideSearchBox && !isClickInsideSuggestions) {
        document.getElementById('suggestions').style.display = 'none';
    }
});

// Function to perform a search and display suggestions based on user input
function searchInfos() {
    let input = document.getElementById('search-input').value;
    if (input.length > 0) {
        let suggestions = [];

        // Fetching data from API for search
        fetch('/api/search/every/informations')
            .then(response => response.json())
            .then(data => {
                // Iterating through fetched data to find matching suggestions
                data.forEach(band => {
                    if (band.name.toLowerCase().includes(input.toLowerCase())) {
                        suggestions.push({
                            name: band.name,
                            type: 'Groupe',
                            image: band.image,
                            redirectTo: band.id
                        });
                    }

                    band.members.forEach(member => {
                        if (member.toLowerCase().includes(input.toLowerCase())) {
                            suggestions.push({
                                name: member,
                                type: 'Membre',
                                image: band.image,
                                redirectTo: band.id
                            });
                        }
                    });

                    if (band.firstAlbum && band.firstAlbum.toLowerCase().includes(input.toLowerCase())) {
                        suggestions.push({
                            name: band.firstAlbum,
                            type: 'FirstAlbum',
                            image: band.image,
                            redirectTo: band.id
                        });
                    }

                    if (band.creationDate && band.creationDate.toString().toLowerCase().includes(input.toLowerCase())) {
                        suggestions.push({
                            name: band.creationDate.toString(),
                            type: 'CreationDate',
                            image: band.image,
                            redirectTo: band.id
                        });
                    }

                    band.locations.forEach(location => {
                        if (location.toLowerCase().includes(input.toLowerCase())) {
                            suggestions.push({
                                name: location,
                                type: 'Location',
                                image: band.image,
                                redirectTo: band.id,
                            });
                        }
                    });
                });

                // Displaying the suggestions
                showSuggestions(suggestions);
            })
            .catch(error => console.error('Error:', error));
    } else {
        document.getElementById('suggestions').style.display = 'none';
    }
}

// Function to display search suggestions
function showSuggestions(suggestions) {
    let suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';

    if (suggestions.length > 0) {
        // Adding suggestion elements to the suggestions container
        suggestions.forEach(suggestion => {
            let suggestionElement = document.createElement('div');
            suggestionElement.classList.add('suggestion-item');

            if (suggestion.image) {
                let suggestionImage = document.createElement('img');
                suggestionImage.src = suggestion.image;
                suggestionImage.alt = suggestion.name;
                suggestionImage.classList.add('artist-image');
                suggestionElement.appendChild(suggestionImage);
            }

            let suggestionName = document.createElement('p');
            suggestionName.classList.add('name-artist');

            // Constructing suggestion text based on type
            if (suggestion.type === "Membre") {
                suggestionName.textContent = suggestion.name + " (Membre)";
            } else if (suggestion.type === "Groupe") {
                suggestionName.textContent = suggestion.name + " (Groupe)";
            } else if (suggestion.type === "FirstAlbum") {
                suggestionName.textContent = suggestion.name + " (Premier Album)";
            } else if (suggestion.type === "CreationDate") {
                suggestionName.textContent = suggestion.name + " (Date de Création)";
            } else if (suggestion.type === "Location") {
                suggestionName.textContent = suggestion.name + " (Lieu)";
            }

            suggestionElement.appendChild(suggestionName);

            // Adding click event listener to redirect user on suggestion click
            suggestionElement.addEventListener('click', function() {
                let redirectId = suggestion.redirectTo ? suggestion.redirectTo : suggestion.id;
                window.location.href = "/artists.html?id=" + redirectId;
            });

            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}
