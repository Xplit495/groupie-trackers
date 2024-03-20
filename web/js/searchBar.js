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
        fetch('/api/search/artists')
            .then(response => response.json())
            .then(data => {
                let suggestions = [];

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

                showSuggestions(suggestions.unique());
            })
            .catch(error => console.error('Error:', error));
    } else {
        document.getElementById('suggestions').style.display = 'none';
    }
}

// Fonction pour filtrer les suggestions uniques
Array.prototype.unique = function() {
    return this.filter(function (value, index, self) {
        return self.indexOf(value) === index;
    });
}

function showSuggestions(suggestions) {
    let suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';

    if (suggestions.length > 0) {
        suggestions.forEach(suggestion => {
            let suggestionElement = document.createElement('div');
            suggestionElement.classList.add('suggestion-item');

            if (suggestion.image) {
                let bandImage = document.createElement('img');
                bandImage.src = suggestion.image;
                bandImage.alt = suggestion.name;
                bandImage.classList.add('artist-image');
                suggestionElement.appendChild(bandImage);
            }

            let bandName = document.createElement('p');
            bandName.classList.add('name-artist');

            if (suggestion.type === "Membre") {
                bandName.textContent = suggestion.name + " (Membre)";
            } else if (suggestion.type === "Groupe") {
                bandName.textContent = suggestion.name + " (Groupe)";
            } else if (suggestion.type === "FirstAlbum") {
                bandName.textContent = suggestion.name + " (Premier Album)";
            } else if (suggestion.type === "CreationDate") {
                bandName.textContent = suggestion.name + " (Date de Création)";
            }

            suggestionElement.appendChild(bandName);

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