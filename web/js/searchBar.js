document.addEventListener('click', function(event) {
    let isClickInsideSearchBox = document.getElementById('search-box').contains(event.target);
    let isClickInsideSuggestions = document.getElementById('suggestions').contains(event.target);

    if (!isClickInsideSearchBox && !isClickInsideSuggestions) {
        document.getElementById('suggestions').style.display = 'none';
    }
});

function searchInfos() {
    let input = document.getElementById('search-input').value;
    if (input.length > 0) {
        let suggestions = [];

        Promise.all([
            fetch('/api/search/locations/search/bar').then(response => response.json()),
            fetch('/api/search/artists').then(response => response.json())
        ])
            .then(([data2, data]) => {
                data2.forEach(item => {
                    item.locations.forEach(location => {
                        if (location.toLowerCase().includes(input.toLowerCase())) {
                            suggestions.push({
                                name: location,
                                type: 'Location',
                                image: item.image,
                                redirectTo: item.id,
                            });
                        }
                    });
                });

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
                });

                showSuggestions(suggestions);
            })
            .catch(error => console.error('Error:', error));
    } else {
        document.getElementById('suggestions').style.display = 'none';
    }
}

function showSuggestions(suggestions) {
    console.log(suggestions);
    console.log(suggestions.length)
    let suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';

    if (suggestions.length > 0) {
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

            if (suggestion.type === "Membre") {
                suggestionName.textContent = suggestion.name + " (Membre)";
            } else if (suggestion.type === "Groupe") {
                suggestionName.textContent = suggestion.name + " (Groupe)";
            } else if (suggestion.type === "FirstAlbum") {
                suggestionName.textContent = suggestion.name + " (Premier Album)";
            } else if (suggestion.type === "CreationDate") {
                suggestionName.textContent = suggestion.name + " (Date de Création)";
            }else if (suggestion.type === "Location") {
                suggestionName.textContent = suggestion.name + " (Lieu)";
            }

            suggestionElement.appendChild(suggestionName);

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