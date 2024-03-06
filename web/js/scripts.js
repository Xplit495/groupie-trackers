
function searchArtist() {
    var input = document.getElementById('search-input').value;
    if (input.length > 0) {
        fetch('https://groupietrackers.herokuapp.com/api/artists')
            .then(response => response.json())
            .then(data => {
                var suggestions = data.filter(artist => artist.Name.toLowerCase().includes(input.toLowerCase()));
                showSuggestions(suggestions);
            })
            .catch(error => console.error('Error:', error));
    } else {
        document.getElementById('suggestions').innerHTML = '';
    }
}

function showSuggestions(suggestions) {
    var suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = ''; // Nettoyer les suggestions précédentes
    suggestions.forEach(artist => {
        var suggestionElement = document.createElement('div');
        suggestionElement.innerHTML = `<a href="/artist-details.html?id=${artist.ID}">${artist.Name}</a>`;
        suggestionsContainer.appendChild(suggestionElement);
    });
}